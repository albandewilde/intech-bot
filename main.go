package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/albandewilde/intech-bot/discordclient"
	"github.com/albandewilde/intech-bot/metrics"
)

var (
	TKN  string // Discord bot token
	HOST string // Metric host server
	PORT int64  // Metric port server
)

func init() {
	// Get discord token
	TKN = os.Getenv("TKN")
	if TKN == "" {
		log.Fatal("No discord token provided in the environment variable `TKN`")
	}

	// Get host and port
	HOST = os.Getenv("HOST")
	if HOST == "" {
		HOST = "0.0.0.0"
	}
	port := os.Getenv("PORT")
	var err error
	PORT, err = strconv.ParseInt(port, 10, 64)
	if err != nil {
		PORT = 5419
	}
}

func main() {
	// Start the discord bot
	bot, err := discordclient.NewInitializedBot(TKN)
	if err != nil {
		log.Fatal(err)
	}

	// Start the metric server
	srv := metrics.NewStartedMetricServer("0.0.0.0", 5419)

	// Gracefully close the discord bot and the metric server
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	errs := discordclient.Close(bot)
	if len(errs) > 0 {
		log.Fatal(errs)
	}
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalln(err)
	}

}
