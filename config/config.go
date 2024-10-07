package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botPrefix"`
	config    *Config
}

var (
	Token     string
	BotPrefix string
)

func ReadConfig() error {
	fmt.Println("Reading the config.json")
	file, err := os.ReadFile("./config.json")
	if err != nil {
		log.Panic("There has been an error reading the file")
		os.Exit(1)
	}

	fmt.Println("Unmarshaling the json file")

	var botConfig Config
	err = json.Unmarshal(file, &botConfig)

	if err != nil {
		fmt.Println("There is a problem with the json")
		os.Exit(1)
	}
	Token = botConfig.Token
	BotPrefix = botConfig.BotPrefix

	return nil

}
