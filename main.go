package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
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
	bot, err := NewBot(TKN)
	if err != nil {
		log.Fatal(err)
	}

	err = bot.Open()
	if err != nil {
		log.Fatal(err)
	}

	RegisterCommands(bot)

	// Gracefully close the discord bot
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	bot.Close()
}
