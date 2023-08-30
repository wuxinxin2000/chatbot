package routers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	// "github.com/astaxie/beego/validation"
	// "strconv"
	"chatbot/models"
)

// Init router for handling incoming request
func InitRouter() *gin.Engine {
	 r := gin.Default()
	 r.POST("/review", getReview)
	 r.POST("/review/", getReview)
	 r.POST("/followup", followup)
	 r.POST("/followup/", followup)
	return r
}

// data structure for receiving incoming request data
type request_info struct {
	ID         int `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
  Message string `gorm:"not null" json:"message"` 
}

func getReview(c *gin.Context)  {
	fmt.Printf("got /review request\n")
	const waiting_for_feedback string = "Waiting for feedback"
	var req_info request_info
	var customer models.Customer
	var customer_status models.CustomerStatus
	var chat_template models.ChatTemplate
	c.BindJSON(&req_info)
	// retrieve customer name from db by using the given id
	customer = models.GetCustomerName(req_info.ID)	//strconv.Atoi()
	customer_status = models.GetCustomerStatus(req_info.ID)	//strconv.Atoi()
	// valid := validation.Validation{}
	// valid.Required(name, "name").Message("name must be not null")
	// valid.MaxSize(name, 100, "name").Message("name must be within 100 characters")
	// if ! valid.HasErrors() {
	// }

	// retrieve template body message
	if strings.Contains(req_info.Message, "thanks") || strings.Contains(req_info.Message, "thank") {
		chat_template = models.GetChatTemplate("thanks")
	} else if strings.Contains(req_info.Message, "purchase") || strings.Contains(req_info.Message, "buy") || strings.Contains(req_info.Message, "subscribe") {
		chat_template = models.GetChatTemplate("purchase")
	} else if strings.Contains(req_info.Message, "know more") || strings.Contains(req_info.Message, "about product") {
		chat_template = models.GetChatTemplate("product")
	} else if strings.Contains(req_info.Message, "provide feedback") || strings.Contains(req_info.Message, "give feedback") {
		chat_template = models.GetChatTemplate("need_feedback")
		if (models.CustomerStatus{}) != customer_status {
			models.UpdateCustomerStatus(req_info.ID, waiting_for_feedback)
		} else {
			models.InsertCustomerStatus(req_info.ID, waiting_for_feedback)
		}
	} else if customer_status.Status == waiting_for_feedback {
		chat_template = models.GetChatTemplate("received_feedback")
		models.UpdateCustomerStatus(req_info.ID, "New")
	} else {
		chat_template = models.GetChatTemplate("welcome")
	}

	// add review record into db
	returned_message := "Hi " + customer.Name +  ", " + chat_template.TemplateBody
	models.PostReview(customer.ID, chat_template.TemplateID, req_info.Message, returned_message)
  c.IndentedJSON(http.StatusOK, gin.H{"message": returned_message, })

	// data := make(map[string]interface{})
	// data["customer_id"] = req_info.ID
	// data["customer_name"] = customer.Name
	// data["response_message"] = returned_message

  // c.IndentedJSON(http.StatusOK, gin.H{"message": returned_message, 
  //       "data" : data,})
}

func followup(c *gin.Context)  {
	fmt.Printf("got /followup request\n")
	var req_info request_info
	var customer models.Customer
	var chat_template models.ChatTemplate
	c.BindJSON(&req_info)
	// retrieve customer name from db by using the given id
	customer = models.GetCustomerName(req_info.ID)
	chat_template = models.GetChatTemplate("need_feedback")
	// add review record into db
	returned_message := "Hi " + customer.Name +  ", " + chat_template.TemplateBody
	models.PostReview(customer.ID, chat_template.TemplateID, req_info.Message, returned_message)
  c.IndentedJSON(http.StatusOK, gin.H{"message": returned_message, })
}