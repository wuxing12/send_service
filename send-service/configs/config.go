package configs

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type AppConfig struct {
	Sender     string   `json:"sender"`
	Password   string   `json:"password"`
	Host       string   `json:"host"`
	Port       uint     `json:"port"`
	Recipients []string `json:"recipients"`
}

var ac = AppConfig{}

func LoadConfig(configFile string) {
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("fail to read config file: %s, error: %v", configFile, err)
	}
	if err = json.Unmarshal(bytes, &ac); err != nil {
		log.Fatalf("config file decode error: %v", err)
	}
}

func GetConfig() AppConfig {
	return ac
}
