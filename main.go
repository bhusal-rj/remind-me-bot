package main

import (
	"github.com/bhusal-rj/remind-me/bot"
	"github.com/bhusal-rj/remind-me/config"
)

func main() {
	config.ReadConfig()
	bot.Start()
}
