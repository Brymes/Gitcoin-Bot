package bot

import (
	"Brymes-Gitcoin-Bot/config"
	"Brymes-Gitcoin-Bot/utils"
	"github.com/bwmarrin/discordgo"
	"log"
)

func InitBot() {
	buff, logger := config.InitLoggerInstance("Get-Bounties")
	go utils.PeriodicLog(buff)

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + config.DiscordBotToken)
	if err != nil {
		log.Fatalln("error creating Discord session,", err)
	}

	//Ensure Messages only come in from Guilds
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	RegisterHandlers(dg)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		log.Fatal("Error Initializing bot")
	}

	RegisterCommands(dg, logger)

	config.DiscordBot = dg

}
