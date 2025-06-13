package Baro

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

type InventoryItem struct {
	Item    string `json:"item"`
	Ducats  int    `json:"ducats"`
	Credits int    `json:"credits"`
}

func BaroHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	go func() {
		resp, err := http.Get("https://api.warframestat.us/pc/voidTrader")
		if err != nil {
			s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
				Content: "❌ Erreur lors de la récupération des données de Baro.",
			})
			return
		}
		defer resp.Body.Close()

		var baro BaroData
		if err := json.NewDecoder(resp.Body).Decode(&baro); err != nil {
			s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
				Content: "❌ Erreur de décodage des données de Baro.",
			})
			return
		}

		embed := &discordgo.MessageEmbed{
			Title: "Baro Ki'Teer",
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://static.wikia.nocookie.net/warframe/images/2/24/Baro_Ki%27teer.jpg/revision/latest?cb=20200713143414&path-prefix=fr"},
			Color:     0x3498db,
			Timestamp: time.Now().Format(time.RFC3339),
		}

		if baro.Active {
			embed.Description = fmt.Sprintf("Baro est **présent** à `%s` !", baro.Location)
			embed.Fields = []*discordgo.MessageEmbedField{
				{
					Name:   "Temps restant",
					Value:  time.Until(baro.Expiry).Truncate(time.Minute).String(),
					Inline: true,
				},
			}
		} else {
			embed.Description = fmt.Sprintf("Baro **arrive bientôt** à `%s` !", baro.Location)
			embed.Fields = []*discordgo.MessageEmbedField{
				{
					Name:   "Arrive dans",
					Value:  time.Until(baro.Activation).Truncate(time.Minute).String(),
					Inline: true,
				},
				{
					Name:   "Heure d'arrivée (UTC)",
					Value:  baro.Activation.UTC().Format("15:04 - 02 Jan 2006"),
					Inline: true,
				},
			}
		}

		for _, raw := range baro.Inventory {
			itemJson, _ := json.Marshal(raw)
			var item InventoryItem
			if err := json.Unmarshal(itemJson, &item); err == nil {
				field := &discordgo.MessageEmbedField{
					Name:   fmt.Sprintf("%s", item.Item),
					Value:  fmt.Sprintf("%d ducats | %d crédits", item.Ducats, item.Credits),
					Inline: true,
				}
				embed.Fields = append(embed.Fields, field)
			}
		}

		s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
			Embeds: []*discordgo.MessageEmbed{embed},
		})
	}()
}
