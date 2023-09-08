package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type EnvConfig struct {
	Token            string
	MongoUri         string
	AppId            string
	GuildId          string
	ChannelId        string
	DevChannelId     string
	VoiceChannelId   string
	VoiceChannelLink string
	VtApiKey         string
}

var Env *EnvConfig

func LoadEnv() {
	if Env != nil {
		return
	}

	Env = &EnvConfig{
		Token:            os.Getenv("TOKEN"),
		MongoUri:         os.Getenv("MONGO_URI"),
		AppId:            os.Getenv("APP_ID"),
		GuildId:          os.Getenv("GUILD_ID"),
		ChannelId:        os.Getenv("CHANNEL_ID"),
		DevChannelId:     os.Getenv("DEV_CHANNEL_ID"),
		VoiceChannelId:   os.Getenv("VOICE_CHANNEL_ID"),
		VoiceChannelLink: os.Getenv("VOICE_CHANNEL_LINK"),
		VtApiKey:         os.Getenv("VT_API_KEY"),
	}
}
