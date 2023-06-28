package manager

import (
	"Brymes-Gitcoin-Bot/config"
	"Brymes-Gitcoin-Bot/utils"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func Manager(channel string) {
	if config.ManagerRunning {
		return
	}

	config.ManagerRunning = true
	ticker := time.NewTicker(3 * time.Minute)

	for {
		select {
		case <-ticker.C:
			if config.BountyBotActive {
				go GetBounties(channel)
			}

			if config.GrantBotActive {
				go GetGrants(channel)
			}
		default:
			continue
		}
	}
}

func sendEmbed(channel string, embed *discordgo.MessageEmbed, logger *log.Logger) {
	_, err := config.DiscordBot.ChannelMessageSendEmbed(channel, embed)
	utils.OnErrorPanic(err, "Problem Sending Update", logger)
}
