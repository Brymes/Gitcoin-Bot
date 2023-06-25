package bot

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	RegisteredCommands []*discordgo.ApplicationCommand
	TextChannelType    = []discordgo.ChannelType{
		discordgo.ChannelTypeGuildText,
		discordgo.ChannelTypeGroupDM,
		discordgo.ChannelTypeDM,
		discordgo.ChannelTypeGuildNews,
		discordgo.ChannelTypeGuildNewsThread,
		discordgo.ChannelTypeGuildCategory,
		discordgo.ChannelTypeGuildPrivateThread,
		discordgo.ChannelTypeGuildPublicThread,
		discordgo.ChannelTypeGuildStore,
	}
)

func RegisterCommands(discordSession *discordgo.Session, logger *log.Logger) {
	var commands = []*discordgo.ApplicationCommand{
		{
			Name:        "help",
			Description: "Returns All Commands and Descriptions",
		},
		{
			Name:        "track_grants",
			Description: "Starts Sending GitCoin Grant updates to channel it's called on",
		},
		{
			Name:        "track_bounties",
			Description: "Starts Sending GitCoin Bounty updates to channel it's called on",
		},
	}

	logger.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))

	for index, command := range commands {
		cmd, err := discordSession.ApplicationCommandCreate(discordSession.State.User.ID, "", command)
		if err != nil {
			logger.Fatalf("Cannot create '%v' command: %v", command.Name, err)
		}
		registeredCommands[index] = cmd
		logger.Println("Adding command", cmd.Name)

	}
	RegisteredCommands = registeredCommands
}

func DeRegisterCommands(discordSession *discordgo.Session, logger *log.Logger) {
	logger.Println("Removing commands...")
	// We need to fetch the commands, since deleting requires the command ID.
	// We are doing this from the returned commands on line 375, because using
	// this will delete all the commands, which might not be desirable, so we
	// are deleting only the commands that we added.
	// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
	// if err != nil {
	// 	logger.Fatalf("Could not fetch registered commands: %v", err)
	// }

	for _, command := range RegisteredCommands {
		err := discordSession.ApplicationCommandDelete(discordSession.State.User.ID, "", command.ID)
		if err != nil {
			logger.Panicf("Cannot delete '%v' command: %v", command.Name, err)
		}
	}
}
