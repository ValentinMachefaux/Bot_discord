package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"net/http"
	"os"
)

//func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
//	if m.Author.ID == BotId {
//		return
//	}
//
//	if m.Content == BotPrefix+"hello" {
//		_, _ = s.ChannelMessageSend(m.ChannelID, "Ordis vous dit bonjour !")
//	}
//
//	if m.Content == BotPrefix+"baro" {
//		resp, err := http.Get("https://api.warframestat.us/pc/voidTrader")
//		if err != nil {
//			fmt.Println("No response from request")
//		}
//		defer func(Body io.ReadCloser) {
//			err := Body.Close()
//			if err != nil {
//				fmt.Println(err)
//			}
//		}(resp.Body)
//		body, err := io.ReadAll(resp.Body)
//
//		var B bin.BaroJson
//		if err := json.Unmarshal(body, &B); err != nil {
//			fmt.Println("Can not unmarshal JSON")
//		}
//		if B.Active == false {
//			//fmt.Println(B.Character + "\n" + B.Location)
//			//fmt.Println(B.Inventory)
//			embed := &discordgo.MessageEmbed{
//				Author: &discordgo.MessageEmbedAuthor{},
//				Title:  B.Character,
//				Color:  0x7a00ff,
//				Thumbnail: &discordgo.MessageEmbedThumbnail{
//					URL: "https://n9e5v4d8.ssl.hwcdn.net/uploads/663c26145b1654c2264392a09568b5c2.png",
//				},
//				Fields: []*discordgo.MessageEmbedField{
//					{
//						Name: "~~",
//					},
//					{
//						Name:  "Location:",
//						Value: B.Location,
//					},
//					{
//						Name:   "Start in :",
//						Value:  B.StartString,
//						Inline: true,
//					},
//					{
//						Name:   "End in :",
//						Value:  B.EndString,
//						Inline: true,
//					},
//					{
//						Name: "~~",
//					},
//					{
//						Name: "Inventory : ",
//					},
//					{
//						Name:   "Start in :",
//						Value:  B.StartString,
//						Inline: true,
//					},
//				},
//			}
//			sendEmbed, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
//			if err != nil {
//				return
//			}
//		} else {
//			fmt.Println("Baro Ki'teer got problem on the road")
//		}
//	}
//}

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
