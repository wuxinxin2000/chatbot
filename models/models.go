package models

import (
	"chatbot/setting"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
}

type Customer struct {
	Model
	Name       string `gorm:"not null" json:"name"`
	Email  		 string `json:"email"`
	Gender 		 string `json:"gender"`
}

type CustomerStatus struct {
	ID         int `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Status     string `json:"status"`
}

type ChatTemplate struct {
	TemplateID		int `gorm:"primary_key;AUTO_INCREMENT" json:"template_id"`
  TemplateName string `gorm:"not null" json:"template_name"` 
  TemplateType string `json:"template_type"`
  TemplateBody string `json:"template_body"`
	CreatedAt  time.Time `json:"created_at"`
}

type Chats struct {
	RecordID		int `gorm:"primary_key;AUTO_INCREMENT" json:"record_id"`
	ChatID		string `gorm:"not null" json:"chat_id"`
  CustomerID	int `gorm:"not null" json:"customer_id"` 
  TemplateID	int `gorm:"not null" json:"template_id"` 
	ReceivedMessage string `json:"received_message"`
  ReturnedMessage string `json:"returned_message"`
	CreatedAt  time.Time `json:"created_at"`
}

func dsn(user string, password string, host string, dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName)
}

// Setup initializes the database instance
func Setup() {
	var err error
	Db, err = gorm.Open(setting.DatabaseSetting.Type, dsn(
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.DBName))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	Db.LogMode(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(20)
	//  Db.DB().SetConnMaxLifetime(time.Minute * 5)
}

func CloseDB() {
	defer Db.Close()
}

// retrieve customer name from given id
func GetCustomerName(id int) (customer Customer) {
	Db.Where("id = ?", id).First(&customer)
	return
}
// retrieve customer status from given id
func GetCustomerStatus(id int) (customerStatus CustomerStatus) {
	Db.Where("id = ?", id).First(&customerStatus)
	return
}

// update customer status for given id
func UpdateCustomerStatus(id int, status string) {
	Db.Model(&CustomerStatus{}).Where("id = ?", id).Update("status", status)
	Db.Model(&CustomerStatus{}).Where("id = ?", id).Update("created_at", time.Now())
	return
}

// insert customer status for given id
func InsertCustomerStatus(id int, status string) {
	customerStatus := CustomerStatus{ID: id, Status: status, CreatedAt: time.Now()}
	Db.Create(&customerStatus)
	return
}

// get chat template for chatbot response message
func GetChatTemplate(template_type string) (chat_template ChatTemplate) {
	Db.Where("template_type = ?", template_type).First(&chat_template)
	return
}

// create and save a review record for the conversation between customer and chatbot
func PostChat(chat_id string, customer_id int, template_id int, received_message string, returned_message string) {
	chat := Chats{ChatID: chat_id, CustomerID: customer_id, TemplateID: template_id, ReceivedMessage: received_message, ReturnedMessage: returned_message, CreatedAt: time.Now()}
	Db.Create(&chat)
}
