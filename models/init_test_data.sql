use chatbotdb;
INSERT INTO customers (name, email, gender) VALUES ('Alice', 'alice@gmail.com', 'F');
INSERT INTO customers (name, email, gender) VALUES ('Bob', 'bob@gmail.com', 'M');
INSERT INTO customers (name, email, gender) VALUES ('Charlie', 'charlie@gmail.com', 'M');
INSERT INTO customers (name, email, gender) VALUES ('Daisy', 'daisy@gmail.com', 'F');


INSERT INTO chat_templates (template_name, template_type, template_body) VALUES ('hello', 'welcome', "Welcome to Connectly.ai! How can I help you?\nYou can ask me something like below:\n1. How to subscribe service from connectly.ai? \n2. I would like to know more about connectly.ai\n3. I would like to provide feedback for connectly.ai\n");
INSERT INTO chat_templates (template_name, template_type, template_body) VALUES ('purchase', 'purchase', 'Please go to website https://www.connectly.ai/pricing to check our price. Thanks!');
INSERT INTO chat_templates (template_name, template_type, template_body) VALUES ('intro', 'product', 'Please go to website https://www.connectly.ai/products-for-businesses to check our product. Thanks!');

INSERT INTO chat_templates (template_name, template_type, template_body) VALUES ('thank_you', 'thanks', 'Thank you for your trust and loyalty to our products!');
INSERT INTO chat_templates (template_name, template_type, template_body) VALUES ('need_feedback', 'need_feedback', 'Please feel free to leave your feedback to help us serve you better. Any feedback is appreciated! Thank you!');
INSERT INTO chat_templates (template_name, template_type, template_body) VALUES ('received_feedback', 'received_feedback', 'Thanks for your precious feedback to our products!');

-- Data in both chats table and customer_statuses should be inserted on the fly when handling the triggered request.
