package commands

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func message(s *discordgo.Session, i *discordgo.Interaction, msg string) {
	res := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	}
	s.InteractionRespond(i, res)
}

func unexpectedError(s *discordgo.Session, i *discordgo.Interaction, err error) {
	slog.Error(err.Error())
	s.InteractionRespond(i, errorResponse("Something went wrong! Try again later."))
}

func userError(s *discordgo.Session, i *discordgo.Interaction, msg string) {
	s.InteractionRespond(i, errorResponse(msg))
}

func errorResponse(msg string) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	}
}
