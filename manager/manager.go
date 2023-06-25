package manager

import (
	"Brymes-Gitcoin-Bot/config"
	"Brymes-Gitcoin-Bot/utils"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func Manager(channel string) {
	ticker := time.NewTicker(6 * time.Hour)

	for {
		select {
		case <-ticker.C:
			if config.BountyBotActive {
				go GetBounties(channel)
			}
			//else if config.GrantBotActive {
			//	go GetGrants(channel)
			//}
		default:
			continue
		}
	}
}

func sendEmbed(channel string, embed *discordgo.MessageEmbed, logger *log.Logger) {
	_, err := config.DiscordBot.ChannelMessageSendEmbed(channel, embed)
	utils.OnErrorPanic(err, "Problem Sending Update", logger)
}
