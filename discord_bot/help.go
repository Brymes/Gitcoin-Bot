package bot

import (
	"Brymes-Gitcoin-Bot/utils"
	"github.com/bwmarrin/discordgo"
	"strings"

	"github.com/olekukonko/tablewriter"
)

var (
	HelpText  = formatHelpText(helpArray[:7])
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
	err := discordSession.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: HelpText,
			Title:   "Help/Commands Message",
		},
	})

	utils.OnErrorPanic(err, "Error sending Help Message", nil)

	err = SendChannelSetupFollowUp(HelpText, discordSession, interaction)
	utils.OnErrorPanic(err, "Error sending Help Follow Up", nil)
}

func SendHelpText(discordSession *discordgo.Session, message *discordgo.MessageCreate) {
	_, err := discordSession.ChannelMessageSend(message.ChannelID, HelpText)
	utils.OnErrorPanic(err, "Error sending Help Message", nil)
}

func SendChannelSetupFollowUp(message string, discordSession *discordgo.Session, interaction *discordgo.InteractionCreate) error {
	_, err := discordSession.FollowupMessageCreate(interaction.Interaction, true, &discordgo.WebhookParams{Content: message})
	return err
}
