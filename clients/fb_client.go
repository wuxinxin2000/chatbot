package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"log"
	"net/http"
	"strings"

	"chatbot/models"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	uriSendMessage = "https://graph.facebook.com/v17.0/me/messages"

	defaultRequestTimeout = 10 * time.Second
)

// https://developers.facebook.com/docs/messenger-platform/send-messages/#messaging_types
const (
	messageTypeResponse = "RESPONSE"
)

var (
	client = &http.Client{}
)

// errors
var (
	errUnknownWebHookObject = errors.New("unknown web hook object")
	errNoMessageEntry       = errors.New("there is no message entry")
)

// HandleMessenger handles all incoming webhooks from Facebook Messenger.
func HandleMessenger(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		fmt.Println("Got http.MethodGet request")
		HandleVerification(c)
		return
	}

	fmt.Println("Calling HandleWebHook")
	HandleWebHook(c)
}

// HandleVerification handles the verification request from Facebook.
func HandleVerification(c *gin.Context) {
	fmt.Println("In HandleVerification")
	if VerifyToken != c.Request.URL.Query().Get("hub.verify_token") || "subscribe" != c.Request.URL.Query().Get("hub.mode") {
		fmt.Printf("VerifyToken should be %v", c.Request.URL.Query().Get("hub.verify_token"))
		c.Writer.WriteHeader(http.StatusUnauthorized)
		c.Writer.Write(nil)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(c.Request.URL.Query().Get("hub.challenge")))
}

// HandleWebHook handles a webhook incoming from Facebook.
func HandleWebHook(c *gin.Context) {
	fmt.Println("In HandleWebHook")
	r := c.Request
	err := Authorize(r)
	if err != nil {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		c.Writer.Write([]byte("unauthorized"))
		log.Println("authorize", err)
		return
	}

	var wr WebHookRequest
	json.NewDecoder(r.Body).Decode(&wr)
	// err = json.Unmarshal(body, &wr)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.Writer.Write([]byte("bad request"))
		log.Println("unmarshal request", err)
		return
	}

	log.Println("WebHookRequest wr: ", wr)
	err = handleWebHookRequest(wr)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("internal"))
		log.Println("handle webhook request", err)
		return
	}

	// Facebook waits for the constant message to get that everything is OK
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte("EVENT_RECEIVED"))
}

func handleWebHookRequest(r WebHookRequest) error {
	if r.Object != "page" {
 		fmt.Printf("WebHookRequest: %v", r)
		return errUnknownWebHookObject
	}

	for _, we := range r.Entry {
		err := handleWebHookRequestEntry(we)
		if err != nil {
			return fmt.Errorf("handle webhook request entry: %w", err)
		}
	}

	return nil
}

func handleWebHookRequestEntry(we WebHookRequestEntry) error {
    // Facebook claims that the arr always contains a single item but we don't trust them :)
	if len(we.Messaging) == 0 { 
		return errNoMessageEntry
	}

	// em := we.Messaging[0]
	for _, em := range we.Messaging {
		// message action
		err := handleMessage(em.Sender.ID, em.Message.Text)
		if err != nil {
			return fmt.Errorf("handle message: %w", err)
		}
	}
	return nil
}

func GetChatTemplate(msg string) string {
	var chat_template models.ChatTemplate
	if strings.Contains(msg, "thanks") || strings.Contains(msg, "thank") {
		chat_template = models.GetChatTemplate("thanks")
	} else if strings.Contains(msg, "purchase") || strings.Contains(msg, "buy") || strings.Contains(msg, "subscribe") {
		chat_template = models.GetChatTemplate("purchase")
	} else if strings.Contains(msg, "know more") || strings.Contains(msg, "about product") {
		chat_template = models.GetChatTemplate("product")
	} else if strings.Contains(msg, "provide feedback") || strings.Contains(msg, "give feedback") {
		chat_template = models.GetChatTemplate("need_feedback")
		// if (models.CustomerStatus{}) != customer_status {	// not empty
		// 	models.UpdateCustomerStatus(req_info.CustomerID, waitingForFeedback)
		// } else {
		// 	models.InsertCustomerStatus(req_info.CustomerID, waitingForFeedback)
		// }
	// } else if customer_status.Status == waitingForFeedback {
	// 	chat_template = models.GetChatTemplate("received_feedback")
	// 	models.UpdateCustomerStatus(req_info.CustomerID, "New")
	} else {
		chat_template = models.GetChatTemplate("welcome")
	}
	return chat_template.TemplateBody;
}

func handleMessage(recipientID, msgText string) error {
	msgText = strings.TrimSpace(msgText)
	fmt.Printf("In handleMessage, recipientID is %v, input message text is: %v", recipientID, msgText)

	var responseText string
	responseText = GetChatTemplate(msgText)

	return Respond(context.TODO(), recipientID, responseText)
}

// Respond responds to a user in FB messenger.
// https://developers.facebook.com/docs/messenger-platform/get-started#step-3--send-the-customer-a-message
func Respond(ctx context.Context, recipientID, responseText string) error {
	// reqURI := fmt.Sprintf("%s?recipient={id:%s}&message={text:'%s'}&messaging_type=%s&access_token=%s",
	// 											uriSendMessage, recipientID, responseText, messageTypeResponse, AccessToken) 
	// reqURI := fmt.Sprintf("%s?access_token=%s",
	// 											uriSendMessage, AccessToken) 
	return callAPI(ctx, uriSendMessage,
		SendMessageRequest{
		MessagingType: messageTypeResponse,
		RecipientID: User{
			ID: recipientID,
		},
		Message: Message{
			Text: responseText,
		},
	})

}

func callAPI(ctx context.Context, reqURI string, reqBody interface{}) error {//
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&reqBody)
	url := fmt.Sprintf("%s?access_token=%s", reqURI, AccessToken)
	req, err := http.NewRequest("POST", url, body)
	// req, err := http.NewRequest("POST", reqURI, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return nil
}

