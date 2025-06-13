package main

import (
	"fmt"
	"github.com/ValentinMachefaux/Bot_discord/bin"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var ordis bin.Ordis
	var err error
	ordis.Config.ReadFile()
	ordis.Session, err = discordgo.New("Bot " + ordis.Config.BotToken)
	ordis.Session.Identify.Intents = discordgo.IntentsAll

	if err != nil {
		fmt.Println(err)
	}

	err = ordis.Session.Open()
	if err != nil {
		fmt.Println(err)
	}

	ordis.InitSlashCommands()
	ordis.InitSlashCommandsHandlers()

	defer ordis.Session.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	if err != nil {
		fmt.Println(err)
	}
}
