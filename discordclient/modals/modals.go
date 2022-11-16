package modals

import (
	"strings"

	dgo "github.com/bwmarrin/discordgo"

	"github.com/albandewilde/intech-bot/discordclient/assignrolemsg"
)

// modalsHandlers
// key → modal name (used as a prefix in his custom ID)
// value → handler function for the modal
var modalsHandlers = map[string]func(*dgo.Session, *dgo.InteractionCreate){
	assignrolemsg.NAME: assignrolemsg.AssignRolesModalHandler,
}

// ModalsHandlers is the handler of modal
func ModalsHandlers(s *dgo.Session, i *dgo.InteractionCreate) {
	ID := i.ModalSubmitData().CustomID

	for name, h := range modalsHandlers {
		if strings.HasPrefix(ID, name) {
			h(s, i)
			return
		}
	}
}
