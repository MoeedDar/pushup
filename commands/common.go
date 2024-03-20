package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	habitOption = discordgo.ApplicationCommandOption{
		Name:        "habit",
		Description: "The habit you want to progress",
		Type:        discordgo.ApplicationCommandOptionString,
		Required:    true,
		Choices:     habitChoices,
	}
	amountOption = discordgo.ApplicationCommandOption{
		Name:        "amount",
		Description: "How much of the habit do you want to complete per interval",
		Type:        discordgo.ApplicationCommandOptionNumber,
		Required:    true,
	}
	habitChoices = []*discordgo.ApplicationCommandOptionChoice{
		{Name: "pushups", Value: "pushups"},
		{Name: "meditation", Value: "meditation"},
	}
	intervalChoices = []*discordgo.ApplicationCommandOptionChoice{
		{Name: "daily", Value: "daily"},
		{Name: "weekly", Value: "weekly"},
		{Name: "monthly", Value: "monthly"},
	}
)
