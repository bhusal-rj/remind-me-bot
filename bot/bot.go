package bot

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bhusal-rj/remind-me/config"
	"github.com/bhusal-rj/remind-me/gemini"
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
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
	TOKEN := config.InitialConfig.Token_ID
	USER_ID := config.InitialConfig.User_ID
	Channel_ID := config.InitialConfig.Channel_ID

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
	_, err = c.AddFunc("0 9,21 * * *", func() {
		count = count + 1
		message := formatForDiscord(gemini.GetInfoFromGemini())
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
func formatForDiscord(response string) string {
	fmt.Println(response)
	response = strings.ReplaceAll(response, "\\n\\n", " ")
	response = strings.ReplaceAll(response, "\\n", " ")
	response = strings.ReplaceAll(response, `"`, "")

	return response + fmt.Sprintf("<@%s>", config.InitialConfig.User_ID)
}
