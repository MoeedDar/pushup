package commands

import (
	"fmt"
	"pushup/model"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var doneOptions = []*discordgo.ApplicationCommandOption{
	&habitOption,
	&amountOption,
}

func done(s *discordgo.Session, e *discordgo.InteractionCreate) {
	i := e.Interaction
	userId := e.Member.User.ID
	habit := e.ApplicationCommandData().Options[0].StringValue()
	amount := e.ApplicationCommandData().Options[1].FloatValue()

	err := model.Complete(userId, habit, float32(amount))
	if err != nil {
		unexpectedError(s, i, err)
		return
	}

	progress, goal, err := model.Progress(userId, habit)
	if err != nil {
		unexpectedError(s, i, err)
		return
	}

	var msg string
	if progress >= goal {
		msg = "Congratulations! You've reached your goal for " + habit + "!"
	} else {
		msg = "You're making excellent progress!\n"
		msg += makeProgressBar(progress, goal)
		msg += fmt.Sprintf("Only %s more to go until you hit your goal!", strconv.FormatFloat(float64(goal-progress), 'f', -1, 32))
	}

	message(s, i, msg)
}

func makeProgressBar(progress, goal float32) string {
	percentage := int(progress / goal * 100)
	bar := "[" + strings.Repeat("=", percentage/5) + strings.Repeat(" ", (100-percentage)/5) + "]\n"
	return bar
}
