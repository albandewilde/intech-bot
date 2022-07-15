package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/albandewilde/intech-bot/discordclient"
)

var TKN string

func init() {
	// Get discord token
	TKN = os.Getenv("TKN")
	if TKN == "" {
		log.Fatal("No discord token provided in the environment variable `TKN`")
	}
}

func main() {
	// Start the discord bot
	bot, err := discordclient.NewInitializedBot(TKN)
	if err != nil {
		log.Fatal(err)
	}

	// Gracefully close the discord bot
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	errs := discordclient.Close(bot)
	if len(errs) > 0 {
		log.Fatal(errs)
	}

}
