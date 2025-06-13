package bin

import (
	"github.com/ValentinMachefaux/Bot_discord/bin/Baro"
	//"github.com/ValentinMachefaux/Bot_discord/bin/ModSearch"
	"github.com/bwmarrin/discordgo"
)

func GetCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "baro",
			Description: "Renvoie les infos sur Baro Ki'teer",
		},
		{
			Name:        "modsearch",
			Description: "BlaBla",
		},
	}
}

func GetHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"baro": Baro.BaroHandler,
		//"modsearch": ModSearch.modSearchHandler,
	}

}
