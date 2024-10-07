package config

import "os"

type BotConfig struct {
	Token_ID   string
	User_ID    string
	Channel_ID string
}

func (b *BotConfig) LoadConfig() {
	b.Channel_ID = os.Getenv("Channel_ID")
	b.Token_ID = os.Getenv("TOKEN")
	b.User_ID = os.Getenv("USER_ID")
}
