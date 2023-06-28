package manager

import (
	"Brymes-Gitcoin-Bot/config"
	"Brymes-Gitcoin-Bot/gitcoin"
	"Brymes-Gitcoin-Bot/utils"
	"time"
)

func GetGrants(channel string) {
	buff, logger := config.InitLoggerInstance("Get-Bounties")
	go utils.PeriodicLog(buff)

	grants := gitcoin.GitcoinService{}.GetGrants(logger)
	for _, grant := range grants {
		sendEmbed(channel, grant, logger)
		time.Sleep(10 * time.Second)
	}

}
