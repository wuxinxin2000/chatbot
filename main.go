package main

import (
	"fmt"
	"log"

	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"chatbot/models"
	"chatbot/routers"
	"chatbot/setting"
)

func init() {
	setting.Setup()
	models.Setup()
}


func main() {
	 routersInit := routers.InitRouter()
	 readTimeout := setting.ServerSetting.ReadTimeout
	 writeTimeout := setting.ServerSetting.WriteTimeout
	 endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	 server := &http.Server {
		Addr:	endPoint,
		Handler: routersInit,
		ReadTimeout: readTimeout,
		WriteTimeout: writeTimeout,
		MaxHeaderBytes: 1 << 20,
	 }
	 
	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
	models.CloseDB()
}