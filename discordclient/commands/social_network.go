package commands

import (
	"log"

	dgo "github.com/bwmarrin/discordgo"
)

var socialNetworkCommand = &dgo.ApplicationCommand{
	Name:        "social-network",
	Description: "Get school social network links",
}

func socialNetworkHandler(s *dgo.Session, i *dgo.InteractionCreate) {
	err := s.InteractionRespond(
		i.Interaction,
		&dgo.InteractionResponse{
			Type: dgo.InteractionResponseChannelMessageWithSource,
			Data: &dgo.InteractionResponseData{
				Components: []dgo.MessageComponent{
					dgo.ActionsRow{
						Components: []dgo.MessageComponent{
							dgo.Button{
								Emoji: dgo.ComponentEmoji{
									Name: "💼",
								},
								Label: "Linkedin",
								Style: dgo.LinkButton,
								URL:   "https://www.linkedin.com/school/in'tech/",
							},
						},
					},
					dgo.ActionsRow{
						Components: []dgo.MessageComponent{
							dgo.Button{
								Emoji: dgo.ComponentEmoji{
									Name: "🐦",
								},
								Label: "Twitter",
								Style: dgo.LinkButton,
								URL:   "https://twitter.com/intechinfo",
							},
						},
					},
					dgo.ActionsRow{
						Components: []dgo.MessageComponent{
							dgo.Button{
								Emoji: dgo.ComponentEmoji{
									Name: "📸",
								},
								Label: "Instagram",
								Style: dgo.LinkButton,
								URL:   "https://www.instagram.com/intech_paris/",
							},
						},
					},
					dgo.ActionsRow{
						Components: []dgo.MessageComponent{
							dgo.Button{
								Emoji: dgo.ComponentEmoji{
									Name: "⚽",
								},
								Label: "Facebook",
								Style: dgo.LinkButton,
								URL:   "https://www.facebook.com/intechinfo",
							},
						},
					},
					dgo.ActionsRow{
						Components: []dgo.MessageComponent{
							dgo.Button{
								Emoji: dgo.ComponentEmoji{
									Name: "🎥",
								},
								Label: "YouTube",
								Style: dgo.LinkButton,
								URL:   "https://www.youtube.com/channel/UCKunybxZNy_LX_NWgewXHLg",
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
