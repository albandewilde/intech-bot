package commands

import (
	"log"

	"github.com/albandewilde/intech-bot/discordclient/assignrolemsg"
	dgo "github.com/bwmarrin/discordgo"
)

// commands register to the bot
var Commands = []*dgo.ApplicationCommand{
	repositoryLinkCommand,
	alumniLinkCommand,
	socialNetworkCommand,
	assignrolemsg.AssignRolesMsg,
}

// commandsHandlers
// key → the command name
// value → the handler function for the command
var commandHandlers = map[string]func(s *dgo.Session, i *dgo.InteractionCreate){
	repositoryLinkCommand.Name:        repositoryLinkHandler,
	alumniLinkCommand.Name:            alumniLinkHandler,
	socialNetworkCommand.Name:         socialNetworkHandler,
	assignrolemsg.AssignRolesMsg.Name: assignrolemsg.AssignRolesCommandHandler,
}

// CommandsHandlers is the handler of slash commands
func CommandsHandlers(s *dgo.Session, i *dgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}

// RegisterCommands to the bot
// `s` must be opened (`s.Open()` must be called before this function)
func RegisterCommands(s *dgo.Session, commands []*dgo.ApplicationCommand) ([]*dgo.ApplicationCommand, []error) {
	errors := make([]error, 0)
	registeredCommands := make([]*dgo.ApplicationCommand, len(commands))
	for i, c := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", c)
		if err != nil {
			errors = append(errors, err)
		} else {
			registeredCommands[i] = cmd
		}
	}

	return registeredCommands, errors
}

// UnregisterCommands to the bot
// `s` must not be closed (`s.Close()` must not be called before this function)
func UnregisterCommands(s *dgo.Session) []error {
	// Fetch registered commands
	commands, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		log.Fatalf("Could not fetch registered commands: %v", err)
	}

	errors := make([]error, 0)
	for _, c := range commands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", c.ID)
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}
