package commands

import (
	"errors"
	"pushup/model"

	"github.com/bwmarrin/discordgo"
)

var (
	goalOptions = []*discordgo.ApplicationCommandOption{
		&habitOption,
		&amountOption,
		{
			Name:        "interval",
			Description: "How often do you want to track this habit",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
			Choices:     intervalChoices,
		},
	}
)

func goal(s *discordgo.Session, e *discordgo.InteractionCreate) {
	i := e.Interaction
	opts := e.ApplicationCommandData().Options
	userId := e.Member.User.ID

	if len(opts) != 3 {
		unexpectedError(s, i, errors.New("discord fucked up"))
		return
	}

	habit := opts[0].StringValue()
	target := opts[1].FloatValue()
	interval := opts[2].StringValue()

	if err := model.UpsertGoal(userId, habit, interval, float32(target)); err != nil {
		unexpectedError(s, i, err)
		return
	}

	message(s, i, "Good luck on your new goal!")
}
