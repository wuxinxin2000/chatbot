package clients

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"chatbot/models"
	"chatbot/setting"

	"github.com/gin-gonic/gin"
)

const waitingForFeedback string = "Waiting for feedback"

// Init router for handling incoming request
func InitRouter() *gin.Engine {
	 r := gin.Default()
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	 gin.SetMode(setting.ServerSetting.RunMode)

   chats := r.Group("/chats")
	 {
		chats.POST("/review", getReview)
		chats.POST("/review/", getReview)
		chats.POST("/followup", followup)
		chats.POST("/followup/", followup)
	 }
	 r.GET("/webhook", HandleMessenger)
	 r.POST("/webhook", HandleMessenger)
	return r
}


// data structure for receiving incoming request data
type RequestInfo struct {
	CustomerID         int `gorm:"primary_key" json:"customer_id"`	// customer_id
	ChatID	string `gorm:"not null" json:"chat_id"`
  Message string `gorm:"not null" json:"message"` 
}

func getUUID() string {
    newUUID, err := exec.Command("uuidgen").Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Generated UUID: %s\n", newUUID)
		return string(newUUID)
}
func getReview(c *gin.Context)  {
	fmt.Printf("got /review request\n")
	var req_info RequestInfo
	var customer models.Customer
	var customer_status models.CustomerStatus
	var chat_template models.ChatTemplate
	var chat_id string
	c.BindJSON(&req_info)
	// retrieve customer name from db by using the given id
	customer = models.GetCustomerName(req_info.CustomerID)
	customer_status = models.GetCustomerStatus(req_info.CustomerID)
	if req_info.ChatID == "" {
		chat_id = getUUID()
	} else {
		chat_id = req_info.ChatID
	}
	// // Maybe can make use of chat_id to retrieve all msgs from the same conversation for future explore.
  // chat_msgs, err := models.GetChatMessages(req_info.CustomerID, chat_id)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(chat_msgs)

	// retrieve template body message
	if strings.Contains(req_info.Message, "thanks") || strings.Contains(req_info.Message, "thank") {
		chat_template = models.GetChatTemplate("thanks")
	} else if strings.Contains(req_info.Message, "purchase") || strings.Contains(req_info.Message, "buy") || strings.Contains(req_info.Message, "subscribe") {
		chat_template = models.GetChatTemplate("purchase")
	} else if strings.Contains(req_info.Message, "know more") || strings.Contains(req_info.Message, "about product") {
		chat_template = models.GetChatTemplate("product")
	} else if strings.Contains(req_info.Message, "provide feedback") || strings.Contains(req_info.Message, "give feedback") {
		chat_template = models.GetChatTemplate("need_feedback")
		if (models.CustomerStatus{}) != customer_status {	// not empty
			models.UpdateCustomerStatus(req_info.CustomerID, waitingForFeedback)
		} else {
			models.InsertCustomerStatus(req_info.CustomerID, waitingForFeedback)
		}
	} else if customer_status.Status == waitingForFeedback {
		chat_template = models.GetChatTemplate("received_feedback")
		models.UpdateCustomerStatus(req_info.CustomerID, "New")
	} else {
		chat_template = models.GetChatTemplate("welcome")
	}

	// add chat record into db
	returned_message := "Hi " + customer.Name +  ", " + chat_template.TemplateBody
	models.PostChat(chat_id, customer.ID, chat_template.TemplateID, req_info.Message, returned_message)
  c.IndentedJSON(http.StatusOK, gin.H{"message": returned_message, "chat_id": chat_id})

}

func followup(c *gin.Context)  {
	fmt.Printf("got /followup request\n")
	var req_info RequestInfo
	var customer models.Customer
	var chat_template models.ChatTemplate
	var chat_id	string
	c.BindJSON(&req_info)
	// retrieve customer name from db by using the given id
	customer = models.GetCustomerName(req_info.CustomerID)
	chat_template = models.GetChatTemplate("need_feedback")
	if req_info.ChatID == "" {
		chat_id = getUUID()
	} else {
		chat_id = req_info.ChatID
	}
	customer_status := models.GetCustomerStatus(req_info.CustomerID)
	if (models.CustomerStatus{}) != customer_status {	// not empty
    fmt.Printf("customer_status is not empty.\n")
		models.UpdateCustomerStatus(req_info.CustomerID, waitingForFeedback)
	} else {
    fmt.Printf("customer_status is empty.\n")
		models.InsertCustomerStatus(req_info.CustomerID, waitingForFeedback)
	}
	// add review record into db
	returned_message := "Hi " + customer.Name +  ", " + chat_template.TemplateBody
	models.PostChat(chat_id, customer.ID, chat_template.TemplateID, req_info.Message, returned_message)
  c.IndentedJSON(http.StatusOK, gin.H{"message": returned_message, "chat_id": chat_id })
}