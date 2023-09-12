package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)


var cfg *ini.File

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	DBName			string
}

var DatabaseSetting = &Database{}

func Setup(filename string) {
	var err error
	cfg, err = ini.Load(filename)
	if err != nil {
		log.Fatalf("settting.Setup, fail to parse 'setting/chatbot.ini': %v", err)
	}
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}