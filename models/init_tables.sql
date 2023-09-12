create database if not exists chatbotdb;
use chatbotdb;
drop table if exists customers, customer_statuses, chat_templates, chats, reviews, customer, customer_statuss, customer_infos;
create table if not exists customers (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
  name VARCHAR(100) NOT NULL, 
  email VARCHAR(100), 
  gender VARCHAR(1), 
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
create table if not exists customer_statuses (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
  status VARCHAR(100), 
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
create table if not exists chat_templates (
  template_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
  template_name VARCHAR(100), 
  template_type VARCHAR(100),
  template_body TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
create table if not exists chats (
  record_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  chat_id VARCHAR(36) NOT NULL,
  customer_id INT NOT NULL,
  template_id INT NOT NULL,
  received_message TEXT,
  returned_message TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)
