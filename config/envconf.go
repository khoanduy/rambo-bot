package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConf struct {
	Token         string
	MongoUri      string
	AppId         string
	GuildId       string
	ChannelId     string
	DevChanId     string
	VoiceChanId   string
	VoiceChanLink string
	VtApiKey      string
}

var Env *EnvConf

func LoadConf() {
	if Env != nil {
		return
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	Env = &EnvConf{
		Token:         os.Getenv("TOKEN"),
		MongoUri:      os.Getenv("MONGO_URI"),
		AppId:         os.Getenv("APP_ID"),
		GuildId:       os.Getenv("GUILD_ID"),
		ChannelId:     os.Getenv("CHANNEL_ID"),
		DevChanId:     os.Getenv("DEV_CHAN_ID"),
		VoiceChanId:   os.Getenv("VOICE_CHAN_ID"),
		VoiceChanLink: os.Getenv("VOICE_CHAN_LINK"),
		VtApiKey:      os.Getenv("VT_API_KEY"),
	}
}
