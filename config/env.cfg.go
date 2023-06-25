package config

import (
	"Brymes-Gitcoin-Bot/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DiscordBotToken string
	SecretKey       string

	BountyBotActive, GrantBotActive bool

	DiscordBot *discordgo.Session

	MessageFooter = discordgo.MessageEmbedFooter{
		Text:    "Powered by Gitcoin",
		IconURL: "https://bounties.gitcoin.co",
	}
)

func LoadEnv() {
	_ = godotenv.Load()
	SecretKey = os.Getenv("JWT_SECRET_KEY")
	if SecretKey == "" {
		SecretKey = utils.GenerateUniqueID(20, "alphanumeric")
	}

	InitDiscordBot()
}

func InitDiscordBot() {
	DiscordBotToken = os.Getenv("DISCORD_BOT_TOKEN")
	if DiscordBotToken == "" {
		log.Fatalln("DISCORD_BOT_TOKEN environment variable not set")
	}
}
