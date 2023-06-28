package bot

import (
	"Brymes-Gitcoin-Bot/config"
	"Brymes-Gitcoin-Bot/manager"
	"Brymes-Gitcoin-Bot/utils"
	"github.com/bwmarrin/discordgo"
	"log"
)

func TrackBountyHandler(discordSession *discordgo.Session, interaction *discordgo.InteractionCreate) {

	channel := interaction.ChannelID

	//Respond Channel is being Setup
	message := "Start Sending Bounty Updates here"
	err := InteractionResponse(message, "Bounty Tracker", discordSession, interaction)
	utils.OnErrorPanic(err, "Error Responding to Bounty request", nil)

	log.Println(err)

	config.BountyBotActive = true

	go manager.Manager(channel)
}
