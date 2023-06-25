package bot

import (
	"Brymes-Gitcoin-Bot/manager"
	"Brymes-Gitcoin-Bot/utils"
	"github.com/bwmarrin/discordgo"
)

func TrackBountyHandler(discordSession *discordgo.Session, interaction *discordgo.InteractionCreate) {

	channel := interaction.ChannelID

	//Respond Channel is being Setup
	message := "Start Sending Bounty Updates here"
	err := SendChannelSetupFollowUp(message, discordSession, interaction)
	utils.OnErrorPanic(err, "Error sending Bounty Follow Up", nil)

	go manager.Manager(channel)
}
