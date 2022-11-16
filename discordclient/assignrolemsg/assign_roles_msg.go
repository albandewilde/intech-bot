package assignrolemsg

import (
	"encoding/json"
	"fmt"
	"log"

	dgo "github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

const NAME = "assign-roles-msg"

// sessings represent the interaction ID as key and the channel that the message will be writen
var sessions map[string]*dgo.Channel = make(map[string]*dgo.Channel)

var AssignRolesMsg = &dgo.ApplicationCommand{
	Name:        NAME,
	Description: "Create a message for roles assignation",
	Options: []*dgo.ApplicationCommandOption{
		{
			Type:        dgo.ApplicationCommandOptionChannel,
			Name:        "channel",
			Description: "Channel the message will be send in",
			ChannelTypes: []dgo.ChannelType{
				dgo.ChannelTypeGuildText,
			},
			Required: true,
		},
	},
}

func AssignRolesCommandHandler(s *dgo.Session, i *dgo.InteractionCreate) {
	opts := make(map[string]*dgo.ApplicationCommandInteractionDataOption, len(i.ApplicationCommandData().Options))
	for _, opt := range i.ApplicationCommandData().Options {
		opts[opt.Name] = opt
	}

	// Create ID and get channel from parameter
	ID := fmt.Sprintf("%s;%s;%s", NAME, i.Interaction.Member.User.ID, uuid.NewString())
	channel := opts["channel"].ChannelValue(s)

	// Register session
	sessions[ID] = channel

	err := s.InteractionRespond(
		i.Interaction,
		&dgo.InteractionResponse{
			Type: dgo.InteractionResponseModal,
			Data: &dgo.InteractionResponseData{
				CustomID: ID,
				Title:    "Write your message",
				Components: []dgo.MessageComponent{
					dgo.ActionsRow{
						Components: []dgo.MessageComponent{
							dgo.TextInput{
								CustomID:    "title",
								Label:       "Title",
								Style:       dgo.TextInputShort,
								Placeholder: "Rules",
								Required:    true,
								MaxLength:   100,
								MinLength:   0,
							},
						},
					},
					dgo.ActionsRow{
						Components: []dgo.MessageComponent{
							dgo.TextInput{
								CustomID:    "msg",
								Label:       "Message content",
								Style:       dgo.TextInputParagraph,
								Placeholder: "1. Don't bother people\n2. Be gentle",
								Required:    true,
								MaxLength:   2000,
								MinLength:   0,
							},
						},
					},
					dgo.ActionsRow{
						Components: []dgo.MessageComponent{
							dgo.TextInput{
								CustomID:    "text_and_roles",
								Label:       "Roles (Json formated key value pair)",
								Style:       dgo.TextInputParagraph,
								Placeholder: fmt.Sprintf("%s\n%s", "Key is the text, value the role name.", `{"Accept": "rules_ok"}`),
								Required:    true,
								MaxLength:   4000, // Max length accepted by the discord api
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

// TODO: factorise this function
func AssignRolesModalHandler(s *dgo.Session, i *dgo.InteractionCreate) {
	data := i.ModalSubmitData()
	ID := data.CustomID
	title := data.Components[0].(*dgo.ActionsRow).Components[0].(*dgo.TextInput).Value
	msgContent := data.Components[1].(*dgo.ActionsRow).Components[0].(*dgo.TextInput).Value
	jsnRoles := data.Components[2].(*dgo.ActionsRow).Components[0].(*dgo.TextInput).Value

	// Get channel which we need to write the message
	channel, ok := sessions[ID]
	delete(sessions, ID)
	if !ok {
		err := s.InteractionRespond(i.Interaction, &dgo.InteractionResponse{
			Type: dgo.InteractionResponseChannelMessageWithSource,
			Data: &dgo.InteractionResponseData{
				Content: "err: can't find session ID",
				Flags:   uint64(dgo.MessageFlagsEphemeral),
			},
		})
		if err != nil {
			log.Println(err)
		}
		return
	}

	// Validate json roles
	var roles map[string]string
	err := json.Unmarshal([]byte(jsnRoles), &roles)
	if err != nil {
		err := s.InteractionRespond(i.Interaction, &dgo.InteractionResponse{
			Type: dgo.InteractionResponseChannelMessageWithSource,
			Data: &dgo.InteractionResponseData{
				Content: "Invalid json",
				Flags:   uint64(dgo.MessageFlagsEphemeral),
			},
		})
		if err != nil {
			log.Println(err)
		}
		return
	}

	// Find roles (discord object) from theirs name
	// Index roles at each commands to always be up to date
	rs, err := s.GuildRoles(i.GuildID)
	if err != nil {
		err := s.InteractionRespond(i.Interaction, &dgo.InteractionResponse{
			Type: dgo.InteractionResponseChannelMessageWithSource,
			Data: &dgo.InteractionResponseData{
				Content: "err: can't get guild's roles",
				Flags:   uint64(dgo.MessageFlagsEphemeral),
			},
		})
		if err != nil {
			log.Println(err)
		}
		return
	}

	guildRoles := make(map[string]*dgo.Role)
	for _, r := range rs {
		guildRoles[r.Name] = r
	}

	textAndRoles := make(map[string]*dgo.Role)
	for text, roleNameWanted := range roles {
		// Check if role is present in guild
		if r, ok := guildRoles[roleNameWanted]; ok {
			textAndRoles[text] = r
		} else {
			err := s.InteractionRespond(i.Interaction, &dgo.InteractionResponse{
				Type: dgo.InteractionResponseChannelMessageWithSource,
				Data: &dgo.InteractionResponseData{
					Content: fmt.Sprintf(`err: role: "%s", not found in guild's roles`, roleNameWanted),
					Flags:   uint64(dgo.MessageFlagsEphemeral),
				},
			})
			if err != nil {
				log.Println(err)
			}
			return
		}
	}

	// Write message in the channel
	_, err = s.ChannelMessageSendEmbed(
		channel.ID,
		&dgo.MessageEmbed{
			Type:        dgo.EmbedTypeRich,
			Title:       title,
			Description: msgContent,
		},
	)

	if err != nil {
		log.Println(err)
		err := s.InteractionRespond(i.Interaction, &dgo.InteractionResponse{
			Type: dgo.InteractionResponseChannelMessageWithSource,
			Data: &dgo.InteractionResponseData{
				Content: fmt.Sprintf("err: can't send message in channel `%s`", channel.Name),
				Flags:   uint64(dgo.MessageFlagsEphemeral),
			},
		})
		if err != nil {
			log.Println(err)
		}
		return
	}

	// Response the user that it's successfull proceed
	err = s.InteractionRespond(i.Interaction, &dgo.InteractionResponse{
		Type: dgo.InteractionResponseChannelMessageWithSource,
		Data: &dgo.InteractionResponseData{
			Content: msgContent + " " + jsnRoles + " " + ID + " ",
		},
	})
	if err != nil {
		log.Println(err)
	}
}
