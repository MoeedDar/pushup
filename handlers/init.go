package handlers

import (
	"pushup/commands"
	"pushup/single"

	"github.com/bwmarrin/discordgo"
)

func handleInteractionCreate(s *discordgo.Session, e *discordgo.InteractionCreate) {
	if h, ok := commands.Handlers[e.ApplicationCommandData().Name]; ok {
		h(s, e)
	}
}

func init() {
	single.Discord.AddHandler(handleInteractionCreate)
}
