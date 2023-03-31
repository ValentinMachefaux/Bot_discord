package bin

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type Ordis struct {
	Config  Config
	Session *discordgo.Session
}

func (o *Ordis) InitSlashCommands() {
	for _, v := range GetCommands() {
		_, err := o.Session.ApplicationCommandCreate(o.Session.State.User.ID, "", v)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}

func (o *Ordis) InitSlashCommandsHandlers() {
	o.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := GetHandlers()[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}
