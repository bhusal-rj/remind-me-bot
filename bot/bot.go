package bot

import (
	"fmt"
	"github.com/bhusal-rj/remind-me/config"
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"syscall"
)

var count = 1

func sendMessageToChannel(s *discordgo.Session, channelId, message string) {
	_, err := s.ChannelMessageSend(channelId, message)

	if err != nil {
		fmt.Println("There has been an error sending the messge to the channel", err)
	}
}

func sendReminderToUser(s *discordgo.Session, userId, message string) {
	channel, err := s.UserChannelCreate(userId)
	if err != nil {
		fmt.Println("Error creating the DM channel", err)
		return
	}

	_, err = s.ChannelMessageSend(channel.ID, message)
	if err != nil {
		fmt.Println("Error sending the message", err)
		return
	}
}

func Start(botConfig config.BotConfig) {
	TOKEN := botConfig.Token_ID
	USER_ID := botConfig.User_ID
	Channel_ID := botConfig.Channel_ID

	dg, err := discordgo.New("Bot " + TOKEN)
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
	_, err = c.AddFunc("*/1 * * * *", func() {
		count = count + 1
		message := "Spamming You"
		if len(Channel_ID) > 2 {
			sendMessageToChannel(dg, Channel_ID, message)
		}
		if len(USER_ID) > 2 {
			sendReminderToUser(dg, USER_ID, message)
		}
	})
	if err != nil {
		fmt.Println("Error scheduling the cron job", err)
		return
	}
	c.Start()
	fmt.Println("Bot is now running")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}
