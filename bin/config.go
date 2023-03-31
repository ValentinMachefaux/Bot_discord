package bin

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// Config est la Structure de la configuration du Bot
// Champs obligatoires de la conf :
//   - BotToken [string]
//   - GuildId [string] (optional)
type Config struct {
	BotToken string `yaml:"botToken"`
	GuildId  string `yaml:"guildId"`
}

func (c *Config) ReadFile() {
	file, err := os.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		fmt.Println("cannot unmarshal " + err.Error())
	}

}
