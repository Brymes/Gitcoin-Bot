package config

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DiscordBotToken string

	BountyBotActive, GrantBotActive, ManagerRunning bool

	DiscordBot *discordgo.Session

	MessageFooter = discordgo.MessageEmbedFooter{
		Text:    "Powered by Gitcoin",
		IconURL: "https://bounties.gitcoin.co",
	}
)

func LoadEnv() {
	_ = godotenv.Load()
	InitDiscordBot()
}

func InitDiscordBot() {
	DiscordBotToken = os.Getenv("DISCORD_BOT_TOKEN")
	if DiscordBotToken == "" {
		log.Fatalln("DISCORD_BOT_TOKEN environment variable not set")
	}
}
