package config

import (
	"Brymes-Gitcoin-Bot/utils"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DiscordBotToken string
	SecretKey       string
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
		log.Fatalln("Kindly Pass SendGrid API key as an environment variable named SENDGRID_API_KEY")
	}
}
