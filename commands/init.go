package commands

import (
	"pushup/single"

	"github.com/bwmarrin/discordgo"
)

type Handler = func(*discordgo.Session, *discordgo.InteractionCreate)

var (
	Handlers = map[string]Handler{
		"goal": goal,
		"done": done,
	}
	cmds = []*discordgo.ApplicationCommand{
		{
			Name:        "goal",
			Description: "Set up a habit goal",
			Type:        discordgo.ChatApplicationCommand,
			Options:     goalOptions,
		},
		{
			Name:        "done",
			Description: "Track your progress towards your goal",
			Type:        discordgo.ChatApplicationCommand,
			Options:     doneOptions,
		},
	}
)

func init() {
	_, err := single.Discord.ApplicationCommandBulkOverwrite(single.Discord.State.User.ID, "", cmds)
	if err != nil {
		panic(err)
	}
}
