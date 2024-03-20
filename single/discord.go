package single

import (
	"pushup/globals"

	"github.com/bwmarrin/discordgo"
)

var Discord *discordgo.Session

func init() {
	var err error
	Discord, err = discordgo.New("Bot " + globals.Env["DISCORD_CLIENT_SK"])
	if err != nil {
		panic(err)
	}
	if err := Discord.Open(); err != nil {
		panic(err)
	}
}
