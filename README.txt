1. Setup Database with below 2 commands(password is empty):
$ mysql -uroot -p < ./models/init_tables.sql 
$ mysql -uroot -p < ./models/init_test_data.sql  

2. Start chatbot
Go to chatbot directory, and see main.go in it, then run below command:
$ go run main.go

3. Interact with chatbot by using below curl commands:
$ curl -X POST localhost:8000/review/ -H 'Content-Type: application/json' -d '{"message":"","id":3}'
$ curl -X POST localhost:8000/review/ -H 'Content-Type: application/json' -d '{"message":"I would like to subscribe service from connectly.ai","id":3}'
$ curl -X POST localhost:8000/review/ -H 'Content-Type: application/json' -d '{"message":"I would like to know more about connectly.ai","id":3}'
$ curl -X POST localhost:8000/review/ -H 'Content-Type: application/json' -d '{"message":"I would like to give feedback about your product","id":3}'
$ curl -X POST localhost:8000/review/ -H 'Content-Type: application/json' -d '{"message":"I like your service","id":3}' 
$ curl -X POST localhost:8000/review/ -H 'Content-Type: application/json' -d '{"message":"thanks","id":3}'

$ curl -X POST localhost:8000/followup -H 'Content-Type: application/json' -d '{"message":"","id":2}'


4. code structures:
models directory has the scripts helping setup db, and the code interact with db;
