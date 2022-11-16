package discordclient

import (
	"log"

	"github.com/albandewilde/intech-bot/discordclient/commands"
	"github.com/albandewilde/intech-bot/discordclient/modals"
	dgo "github.com/bwmarrin/discordgo"
)

// NewInitializedBot create and initialize a new discord bot instance
func NewInitializedBot(tkn string) (*dgo.Session, []error) {
	bot, err := dgo.New("Bot " + tkn)
	if err != nil {
		return nil, []error{err}
	}

	// Register callback functions
	bot.AddHandler(ready)
	bot.AddHandler(interactionHandler)

	// Open the bot connection
	err = bot.Open()
	if err != nil {
		return nil, []error{err}
	}

	_, errs := commands.RegisterCommands(bot, commands.Commands)
	if len(errs) > 0 {
		return nil, errs
	}

	return bot, nil
}

// Close properly close the bot (unregister commands, ...)
func Close(s *dgo.Session) []error {
	defer s.Close()

	errors := commands.UnregisterCommands(s)

	return errors
}

func ready(s *dgo.Session, r *dgo.Ready) {
	log.Println("Discord bot ready !")
}

func interactionHandler(s *dgo.Session, i *dgo.InteractionCreate) {
	switch i.Type {
	case dgo.InteractionApplicationCommand:
		commands.CommandsHandlers(s, i)
	case dgo.InteractionModalSubmit:
		modals.ModalsHandlers(s, i)
	}
}
