package commands

import (
	"log"

	dgo "github.com/bwmarrin/discordgo"
)

var alumniLinkCommand = &dgo.ApplicationCommand{
	Name:        "alumni",
	Description: "Get link to the alumni plateform",
}

func alumniLinkHandler(s *dgo.Session, i *dgo.InteractionCreate) {
	err := s.InteractionRespond(
		i.Interaction,
		&dgo.InteractionResponse{
			Type: dgo.InteractionResponseChannelMessageWithSource,
			Data: &dgo.InteractionResponseData{
				Content: "The alumni plateform.",
				Components: []dgo.MessageComponent{
					dgo.ActionsRow{
						Components: []dgo.MessageComponent{
							dgo.Button{
								Emoji: dgo.ComponentEmoji{
									Name: "🎓",
								},
								Label: "Plateform alumni",
								Style: dgo.LinkButton,
								URL:   "https://esiea-alumni.fr/",
							},
						},
					},
				},
			},
		},
	)
	if err != nil {
		log.Println(err)
	}
}
