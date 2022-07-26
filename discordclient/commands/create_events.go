package commands

import (
	dgo "github.com/bwmarrin/discordgo"
)

var createEventsCommand = &dgo.ApplicationCommand{
	Name:        "create-events",
	Description: "Create new event",
}

func createEventsHandler(s *dgo.Session, i *dgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &dgo.InteractionResponse{
		Type: dgo.InteractionResponseModal,
		Data: &dgo.InteractionResponseData{
			CustomID: "create_events_" + i.Interaction.Member.User.ID,
			Title:    "Create event",
			Components: []dgo.MessageComponent{
				dgo.ActionsRow{
					Components: []dgo.MessageComponent{
						dgo.TextInput{
							CustomID:    "Name",
							Label:       "What is the name of the event?",
							Style:       dgo.TextInputShort,
							Placeholder: "Event",
							Required:    true,
							MaxLength:   300,
							MinLength:   1,
						},
					},
				},
				dgo.ActionsRow{
					Components: []dgo.MessageComponent{
						dgo.TextInput{
							CustomID:    "date",
							Label:       "What is the date of the event?",
							Style:       dgo.TextInputShort,
							Placeholder: "Date",
							Required:    true,
							MaxLength:   300,
							MinLength:   1,
						},
					},
				},
				dgo.ActionsRow{
					Components: []dgo.MessageComponent{
						dgo.TextInput{
							CustomID:  "description",
							Label:     "Description",
							Style:     dgo.TextInputParagraph,
							Required:  false,
							MaxLength: 2000,
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
