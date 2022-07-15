package commands

import (
	"log"

	dgo "github.com/bwmarrin/discordgo"
)

var repositoryLinkCommand = &dgo.ApplicationCommand{
	Name:        "repository",
	Description: "Get this bot source code !",
}

func repositoryLinkHandler(s *dgo.Session, i *dgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &dgo.InteractionResponse{
		Type: dgo.InteractionResponseChannelMessageWithSource,
		Data: &dgo.InteractionResponseData{
			Content: "Here is the source code.",
			Components: []dgo.MessageComponent{
				dgo.ActionsRow{
					Components: []dgo.MessageComponent{
						dgo.Button{
							Emoji: dgo.ComponentEmoji{
								Name: "💻",
							},
							Label: "Github",
							Style: dgo.LinkButton,
							URL:   "https://github.com/albandewilde/intech-bot",
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Println(err)
	}
}
