package main

import (
	"Brymes-Gitcoin-Bot/config"
	bot "Brymes-Gitcoin-Bot/discord_bot"
	"Brymes-Gitcoin-Bot/utils"
	"bytes"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	CentralLogger    *log.Logger
	CentralLogBuffer *bytes.Buffer
)

func init() {
	config.LoadEnv()
	CentralLogBuffer, CentralLogger = config.InitLoggerInstance("Discord-Bot")
	bot.InitBot()
}

func main() {
	go utils.PeriodicLog(CentralLogBuffer)

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session .
	defer log.Println(CentralLogBuffer)
	defer utils.HandlePanicMacro(recover(), CentralLogger)
	defer config.DiscordBot.Close()
	defer bot.DeRegisterCommands(config.DiscordBot, CentralLogger)
	defer utils.HandlePanicMacro(config.DiscordBot, CentralLogger)
}
