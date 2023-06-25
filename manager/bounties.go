package manager

import (
	"Brymes-Gitcoin-Bot/config"
	"Brymes-Gitcoin-Bot/gitcoin"
	"Brymes-Gitcoin-Bot/utils"
	"time"
)

func GetBounties(channel string) {
	buff, logger := config.InitLoggerInstance("Get-Bounties")
	go utils.PeriodicLog(buff)

	bounties := gitcoin.GitcoinService{}.GetBounties(logger)
	for _, bounty := range bounties {
		sendEmbed(channel, bounty, logger)
		time.Sleep(10 * time.Second)
	}

}
