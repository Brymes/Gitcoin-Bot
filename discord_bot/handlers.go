package bot

import (
	"Brymes-Gitcoin-Bot/utils"
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	SlashCommands = map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
		"help": HelpHandler,
		//"track_grants":   handlers.SubscriptionsHandler,
		"track_bounties": TrackBountyHandler,
	}
)

func RegisterHandlers(discordSession *discordgo.Session) {
	// Register the slash commands Handler func as a callback for MessageCreate events.
	discordSession.AddHandler(centralCommandHandler)

	// Register the message Handler func as a callback for MessageCreate events.
	discordSession.AddHandler(messageHandler)
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	defer utils.HandlePanicMacro(s, nil)

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "ping":
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
		if err != nil {
			panic(err)
		}
	case "pong":
		_, err := s.ChannelMessageSend(m.ChannelID, "Ping!")
		if err != nil {
			panic(err)
		}
	default:
		log.Println(m.Content)
		SendHelpText(s, m)
	}
}

func centralCommandHandler(discordSession *discordgo.Session, interaction *discordgo.InteractionCreate) {
	switch interaction.Type {
	case discordgo.InteractionApplicationCommand:
		defer utils.HandlePanicMacro(discordSession, nil)

		if handler, ok := SlashCommands[interaction.ApplicationCommandData().Name]; ok {
			handler(discordSession, interaction)
		} else {
			SlashCommands["help"](discordSession, interaction)
		}
	}
}
