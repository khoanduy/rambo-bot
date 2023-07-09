package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Conf struct {
	Token            string
	MongoUri         string
	AppId            string
	GuildId          string
	ChannelId        string
	DevChannelId     string
	VoiceChannelId   string
	VoiceChannelLink string
	VirusTotalApiKey string
}

var BotConf *Conf

func LoadConf() {
	if BotConf != nil {
		return
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	BotConf = &Conf{
		Token:            os.Getenv("TOKEN"),
		MongoUri:         os.Getenv("MONGO_URI"),
		AppId:            os.Getenv("APP_ID"),
		GuildId:          os.Getenv("GUILD_ID"),
		ChannelId:        os.Getenv("CHANNEL_ID"),
		DevChannelId:     os.Getenv("DEV_CHANNEL_ID"),
		VoiceChannelId:   os.Getenv("VOICE_CHANNEL_ID"),
		VoiceChannelLink: os.Getenv("VOICE_CHANNEL_LINK"),
		VirusTotalApiKey: os.Getenv("VIRUS_TOTAL_API_KEY"),
	}
}
