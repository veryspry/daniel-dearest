package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	commandPrefix string
	botID         string
	botName       = "daniel-dearest"
	token         = os.Getenv("DISCORD_BOT_TOKEN")
	connToken     = "Bot " + token
	imgURL        = "https://scontent-ort2-2.xx.fbcdn.net/v/t1.0-9/422577_4547010081099_1418811859_n.jpg?_nc_cat=110&_nc_ht=scontent-ort2-2.xx&oh=dc89fac4cf27aeba1c073fa3c6dd6949&oe=5C7BC511"
)

func main() {
	discord, err := discordgo.New(connToken)
	errCheck("error creating discord session", err)
	bot, err := discord.User("@me")
	errCheck("error retrieving account", err)

	fmt.Println(discord)
	fmt.Println(bot)

	botID = bot.ID

	discord.AddHandler(commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err = discord.UpdateStatus(0, "♥")
		if err != nil {
			fmt.Println("Error attempting to set my status")
		}
		servers := discord.State.Guilds
		fmt.Printf("daniel-bot started %d servers", len(servers))
	})

	err = discord.Open()
	errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	commandPrefix = "!"

	<-make(chan struct{})

}

func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botID || user.Bot {
		return
	}

	sendMsg := parser(message.Content, message.Mentions)

	fmt.Println("USER", message.Mentions[0].Discriminator)

	if sendMsg == true {
		discord.ChannelMessageSend(message.ChannelID, imgURL)
	}

}

func parser(msg string, usrs []*discordgo.User) bool {
	w := strings.Split(msg, " ")

	for i := 0; i < len(usrs); i++ {
		if usrs[i].Username == "Xersule" || usrs[i].Discriminator == "0983" {
			return true
		}
	}

	for i := 0; i < len(w); i++ {
		l := strings.ToLower(w[i])
		if l == "daniel" {
			return true
		}
	}

	return false
}
