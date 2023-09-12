package setting

import (
	"testing"

	"github.com/go-ini/ini"
)

func TestMapTo(t *testing.T) {
	var err error
	cfg, err = ini.Load("./chatbot.ini")
	if err != nil {
		t.Fatalf("Fail to parse 'setting/chatbot.ini': %v", err)
	}
	mapTo("server", ServerSetting)
	if ServerSetting.RunMode != "debug" {
		t.Errorf("Server's RunMode should match the value in chatbot.ini - \"debug\".")
	}
	if ServerSetting.HttpPort != 8000 {
		t.Errorf("Server's HttpPort should match the value in chatbot.ini - 8000.")
	}
	if ServerSetting.ReadTimeout != 60 {
		t.Errorf("Server's ReadTimeout should match the value in chatbot.ini - 60.")
	}
	if ServerSetting.WriteTimeout != 60 {
		t.Errorf("Server's WriteTimeout should match the value in chatbot.ini - 60.")
	}

}

func TestSetup(t *testing.T) {
	Setup("./chatbot.ini")
	if ServerSetting.RunMode != "debug" {
		t.Errorf("Server's RunMode should match the value in chatbot.ini - \"debug\".")
	}
	if ServerSetting.HttpPort != 8000 {
		t.Errorf("Server's HttpPort should match the value in chatbot.ini - 8000.")
	}
	if ServerSetting.ReadTimeout != 60 * 1000 * 1000 * 1000 {
		t.Errorf("Server's ReadTimeout should match the value in chatbot.ini - 60 seconds.")
	}
	if ServerSetting.WriteTimeout != 60 * 1000 * 1000 * 1000 {
		t.Errorf("Server's WriteTimeout should match the value in chatbot.ini - 60 seconds.")
	}
	if DatabaseSetting.Type != "mysql" {
		t.Errorf("Database's Type should match the value in chatbot.ini - \"mysql\".")
	}
	if DatabaseSetting.User != "root" {
		t.Errorf("Database's User should match the value in chatbot.ini - \"root\".")
	}
	if DatabaseSetting.Password != "" {
		t.Errorf("Database's Password should match the value in chatbot.ini - \"\".")
	}
	if DatabaseSetting.Host != "127.0.0.1:3306" {
		t.Errorf("Database's Host should match the value in chatbot.ini - \"127.0.0.1:3306\".")
	}
	if DatabaseSetting.DBName != "chatbotdb" {
		t.Errorf("Database's DBName should match the value in chatbot.ini - \"chatbotdb\".")
	}
}