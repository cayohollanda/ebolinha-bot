package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Token for bot access
const token string = "<YOUR BOT KEY>"

// Prefix of commands
const prefix string = "!"

// BotID is the ID of bot
var BotID string

func main() {
	discord, err := discordgo.New("Bot " + token)
	checkErr("Error on connect with the API", err)

	botUser, err := discord.User("@me")
	checkErr("Erro on select the user", err)

	BotID = botUser.ID

	discord.AddHandler(messageHandler)
	discord.AddHandler(connected)
	discord.AddHandler(whenAddedOnServer)

	err = discord.Open()
	checkErr("Error in turn on the bot", err)

	fmt.Println("The bot is running...")

	<-make(chan struct{})
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author

	if user.ID == BotID || user.Bot {
		return
	}

	if message.Content == prefix+"ping" {
		session.ChannelMessageSend(message.ChannelID, "Pong")
	}

	// fmt.Printf("Message: %+v || From: %s\n", message.Message.Content, user)
}

func connected(session *discordgo.Session, ready *discordgo.Ready) {
	err := session.UpdateStatus(0, "BOT")
	checkErr("error on update status of Bot", err)
}

func whenAddedOnServer(session *discordgo.Session, guildMemberAdd *discordgo.GuildMemberAdd) {
	if session == nil || guildMemberAdd == nil {
		return
	}

	user := guildMemberAdd.User
	if user.ID == BotID || user.Bot {
		return
	}

	st, err := session.UserChannelCreate(user.ID)
	checkErr("error on create channel", err)

	session.ChannelMessageSend(st.ID, "Hi, my name is Ebolinha, i'm here to help you to moderate this server.")
}

func checkErr(message string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", message, err)
		panic(err)
	}
}
