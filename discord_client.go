package main

import (
	"log"

	dgo "github.com/bwmarrin/discordgo"
)

// NewBot create a new discord bot instance
func NewBot(tkn string) (*dgo.Session, error) {
	bot, err := dgo.New("Bot " + tkn)
	if err != nil {
		return nil, err
	}

	// Register callback functions
	bot.AddHandler(ready)
	bot.AddHandler(repositoryLink)
	bot.AddHandler(repositoryLinkCommandFunc)

	return bot, nil
}

func RegisterCommands(bot *dgo.Session) error {
	_, err := bot.ApplicationCommandCreate(bot.State.User.ID, "", repositoryLinkCommand)
	if err != nil {
		return err
	}
	return nil
}

func ready(s *dgo.Session, r *dgo.Ready) {
	log.Println("Discord bot ready !")
}

var repositoryLinkCommand = &dgo.ApplicationCommand{
	Name:        "repository",
	Description: "Get this bot source code !",
}

func repositoryLinkCommandFunc(s *dgo.Session, i *dgo.InteractionCreate) {
	if i.ApplicationCommandData().Name != "repository" {
		return
	}
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

func repositoryLink(s *dgo.Session, m *dgo.MessageCreate) {
	if m.Content != "£repository" {
		return
	}
	_, err := s.ChannelMessageSendComplex(m.ChannelID, &dgo.MessageSend{
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
		Reference: &dgo.MessageReference{MessageID: m.ID, ChannelID: m.ChannelID, GuildID: m.GuildID},
	})
	if err != nil {
		log.Println(err)
	}
}
