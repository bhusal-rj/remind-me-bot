package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

var TOKEN = "YOUR BOT TOKEN"
var USER_ID = "YOUR USER ID"

func Start() {
	dg, err := discordgo.New("Bot" + TOKEN)
	if err != nil {
		fmt.Println("Error creating the discord session", err)
		return
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening the connection", err)
		return
	}

	c := cron.New()

	_, err = c.AddFunc("*/5 * * * *", func() {
		sendReminder(dg, USER_ID, "Hello")
	})
	if err != nil {
		fmt.Println("Error scheduling the cron job", err)
		return
	}
	c.Start()
	fmt.Println("Bot is now running")
}

func sendReminder(s *discordgo.Session, userId, message string) {
	channel, err := s.UserChannelCreate(userId)
	if err != nil {
		fmt.Println("Error creating the DM channel")
		return
	}

	_, err = s.ChannelMessageSend(channel.ID, message)
	if err != nil {
		fmt.Println("Error sending the message")
		return
	}
}
