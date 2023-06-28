package bot

import (
	"Brymes-Gitcoin-Bot/config"
	"Brymes-Gitcoin-Bot/manager"
	"Brymes-Gitcoin-Bot/utils"

	"github.com/bwmarrin/discordgo"
)

func TrackGrantHandler(discordSession *discordgo.Session, interaction *discordgo.InteractionCreate) {

	channel := interaction.ChannelID

	//Respond Channel is being Setup
	message := "Start Sending Grant Updates here"
	err := InteractionResponse(message, "Grants Tracker", discordSession, interaction)
	utils.OnErrorPanic(err, "Error Responding to Grant Interaction", nil)

	config.GrantBotActive = true
	go manager.Manager(channel)
}
