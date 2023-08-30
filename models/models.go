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

type Reviews struct {
	ReviewID		int `gorm:"primary_key;AUTO_INCREMENT" json:"review_id"`
  CustomerID int `gorm:"not null" json:"customer_id"` 
  TemplateID int `gorm:"not null" json:"template_id"` 
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

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer Db.Close()
}

// retrieve customer name from given id
func GetCustomerName(id int) (customer Customer) {
	Db.Where("id = ?", id).First(&customer)
	return
}
// retrieve customer name from given id
func GetCustomerStatus(id int) (customerStatus CustomerStatus) {
	Db.Where("id = ?", id).First(&customerStatus)
	return
}

func UpdateCustomerStatus(id int, status string) {
	Db.Model(&CustomerStatus{}).Where("id = ?", id).Update("status", status)
	Db.Model(&CustomerStatus{}).Where("id = ?", id).Update("created_at", time.Now())
	return
}

func InsertCustomerStatus(id int, status string) {
	customerStatus := CustomerStatus{ID: id, Status: status, CreatedAt: time.Now()}
	Db.Create(&customerStatus)
	return
}

func GetChatTemplate(template_type string) (chat_template ChatTemplate) {
	Db.Where("template_type = ?", template_type).First(&chat_template)
	return
}

func PostReview(customer_id int, template_id int, received_message string, returned_message string) {
	review := Reviews{CustomerID: customer_id, TemplateID: template_id, ReceivedMessage: received_message, ReturnedMessage: returned_message, CreatedAt: time.Now()}
	Db.Create(&review)
}
