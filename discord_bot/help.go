package bot

import (
	"Brymes-Gitcoin-Bot/utils"
	"github.com/bwmarrin/discordgo"
	"strings"

	"github.com/olekukonko/tablewriter"
)

var (
	HelpText  = formatHelpText(helpArray)
	helpArray = [][]string{
		{"/help", "Display help message"},
		{"/track_grants", "Starts Sending GitCoin Grant updates to channel it's called on"},
		{"/track_bounties", "Starts Sending GitCoin Bounty updates to channel it's called on"},
	}
)

// Note: could remove SetBorder and SetRowLine if char count needs to be reduced
func formatHelpText(helpArray [][]string) string {
	str := &strings.Builder{}
	table := tablewriter.NewWriter(str)
	table.SetHeader([]string{"Help", "Message"})
	table.AppendBulk(helpArray)
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetHeaderLine(true)
	table.Render()
	return "```" + str.String() + "```"
}

func HelpHandler(discordSession *discordgo.Session, interaction *discordgo.InteractionCreate) {
	err := InteractionResponse(HelpText, "Help/Commands Message", discordSession, interaction)
	utils.OnErrorPanic(err, "Error sending Help Message", nil)
}

func SendHelpText(discordSession *discordgo.Session, message *discordgo.MessageCreate) {
	_, err := discordSession.ChannelMessageSend(message.ChannelID, HelpText)
	utils.OnErrorPanic(err, "Error sending Help Message", nil)
}

func InteractionResponse(message, title string, discordSession *discordgo.Session, interaction *discordgo.InteractionCreate) error {
	err := discordSession.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Title:   title,
		},
	})
	return err
}
