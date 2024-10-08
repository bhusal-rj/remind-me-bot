package config

import (
	"os"
)

type BotConfig struct {
	Token_ID   string
	User_ID    string
	Channel_ID string
	Gemini_Key string
	PROMPT     string
}

var InitialConfig *BotConfig

func (b *BotConfig) LoadConfig() {
	config := &BotConfig{
		Channel_ID: os.Getenv("Channel_ID"),
		Token_ID:   os.Getenv("TOKEN"),
		User_ID:    os.Getenv("USER_ID"),
		Gemini_Key: os.Getenv("GEMINI_KEY"),
		PROMPT:     os.Getenv("PROMPT"),
	}
	InitialConfig = config
}
