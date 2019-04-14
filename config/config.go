package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	ApiKey  string
	Cx      string
	LogFile string
	Port    int
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}
	Config = ConfigList{
		ApiKey:  cfg.Section("GoogleApiKey").Key("api_key").String(),
		Cx:      cfg.Section("CustomSearch").Key("cx").String(),
		LogFile: cfg.Section("Log").Key("log_file").String(),
		Port:    cfg.Section("Web").Key("port").MustInt(),
	}
}
