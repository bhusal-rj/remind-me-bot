package main

import (
	"fmt"
	"github.com/bhusal-rj/remind-me/bot"
	"github.com/bhusal-rj/remind-me/config"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var botConfig config.BotConfig
	botConfig.LoadConfig()
	fmt.Println("Configuration loaded successfully")
	bot.Start(botConfig)
}
