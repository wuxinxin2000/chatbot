This chatbot implements some simple function to reply customer with predefined messages when receiving certain messages.

1. Setup Database with below 2 commands(password is empty):
$ mysql -uroot -p < ./models/init_tables.sql 
$ mysql -uroot -p < ./models/init_test_data.sql  

2. Start chatbot
Go to chatbot directory, and check whether there is file main.go, then run below command to start the chatbot server:
$ go run main.go

3. Interact with chatbot by using below curl commands:
$ curl -X POST localhost:8000/chats/review/ -H 'Content-Type: application/json' -d '{"message":"","customer_id":3}'
$ curl -X POST localhost:8000/chats/review/ -H 'Content-Type: application/json' -d '{"message":"I would like to subscribe service from connectly.ai","customer_id":3,"chat_id": "USE_THE_chat_id_GOT_FROM_ABOVE_HTTP_REQUEST"}'
$ curl -X POST localhost:8000/chats/review/ -H 'Content-Type: application/json' -d '{"message":"I would like to know more about connectly.ai","customer_id":3,"chat_id": "USE_THE_chat_id_GOT_FROM_ABOVE_HTTP_REQUEST"}'
$ curl -X POST localhost:8000/chats/review/ -H 'Content-Type: application/json' -d '{"message":"I would like to give feedback about your product","customer_id":3,"chat_id": "USE_THE_chat_id_GOT_FROM_ABOVE_HTTP_REQUEST"}'
$ curl -X POST localhost:8000/chats/review/ -H 'Content-Type: application/json' -d '{"message":"I like your service","customer_id":3,"chat_id": "USE_THE_chat_id_GOT_FROM_ABOVE_HTTP_REQUEST"}' 
$ curl -X POST localhost:8000/chats/review/ -H 'Content-Type: application/json' -d '{"message":"thanks","customer_id":3,"chat_id": "USE_THE_chat_id_GOT_FROM_ABOVE_HTTP_REQUEST"}'

$ curl -X POST localhost:8000/chats/followup -H 'Content-Type: application/json' -d '{"message":"","customer_id":2}'

eg:
$ curl -X POST localhost:8000/chats/review/ -H 'Content-Type: application/json' -d '{"message":"","customer_id":3}'
$ curl -X POST localhost:8000/chats/review/ -H 'Content-Type: application/json' -d '{"message":"I would like to subscribe service from connectly.ai","customer_id":3,"chat_id": "B0B80DD1-94E4-4473-8C96-5283B4273BDC"}'
$ curl -X POST localhost:8000/chats/review/ -H 'Content-Type: application/json' -d '{"message":"I would like to know more about connectly.ai","customer_id":3,"chat_id": "B0B80DD1-94E4-4473-8C96-5283B4273BDC"}'
$ curl -X POST localhost:8000/chats/review/ -H 'Content-Type: application/json' -d '{"message":"I would like to give feedback about your product","customer_id":3,"chat_id": "B0B80DD1-94E4-4473-8C96-5283B4273BDC"}'
$ curl -X POST localhost:8000/chats/review/ -H 'Content-Type: application/json' -d '{"message":"I like your service","customer_id":3,"chat_id": "B0B80DD1-94E4-4473-8C96-5283B4273BDC"}' 
$ curl -X POST localhost:8000/chats/review/ -H 'Content-Type: application/json' -d '{"message":"thanks","customer_id":3,"chat_id": "B0B80DD1-94E4-4473-8C96-5283B4273BDC"}'

$ curl -X POST localhost:8000/chats/followup -H 'Content-Type: application/json' -d '{"message":"","customer_id":2}'


4. code structures:
models/:  includes the scripts helping setup db, and the code interact with db;
routers/: includes the logic for handling the http requests
setting/: includes a chatbot.ini file saving db and http server settings, a setting.go file having the logic to load .ini setting
main.go:  contains the main logic of starting a chatbot server

