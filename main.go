package main

import (
	"fmt"

	"github.com/bhusal-rj/remind-me/bot"
	"github.com/bhusal-rj/remind-me/config"
	"github.com/bhusal-rj/remind-me/gemini"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var botConfig config.BotConfig
	botConfig.LoadConfig()
	gemini.GetInfoFromGemini()
	fmt.Println("Configuration loaded successfully")
	bot.Start(botConfig)
}
