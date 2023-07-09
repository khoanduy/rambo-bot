package core

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
	"github.com/samothrakii/rambo-bot/conf"
	"github.com/samothrakii/rambo-bot/events"
	"github.com/samothrakii/rambo-bot/slash"
)

func BotInit() *discordgo.Session {
	bot, err := discordgo.New("Bot " + conf.BotConf.Token)
	if err != nil {
		log.Fatal("Error creating Discord session", err)
	}

	return bot
}

func RegisterMessageReplyBehavior(bot *discordgo.Session) {
	log.Println("Register message reply behavior...")

	bot.AddHandler(MessageReply)
}

func RegisterSlashCommands(bot *discordgo.Session) []*discordgo.ApplicationCommand {
	log.Println("Registering commands...")

	bot.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := slash.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	regdCmds := make([]*discordgo.ApplicationCommand, len(slash.Commands))
	for i, v := range slash.Commands {
		cmd, err := bot.ApplicationCommandCreate(conf.BotConf.AppId, conf.BotConf.GuildId, v)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", v.Name, err)
		}
		regdCmds[i] = cmd
	}

	return regdCmds
}

func RemoveSlashCommands(bot *discordgo.Session, regdCmds []*discordgo.ApplicationCommand) {
	log.Println("Removing commands...")

	for _, v := range regdCmds {
		err := bot.ApplicationCommandDelete(conf.BotConf.AppId, conf.BotConf.GuildId, v.ID)
		if err != nil {
			log.Fatalf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}

func GameNightEventCron(bot *discordgo.Session) *cron.Cron {
	fmt.Println("Start Game night event cron job")
	job := cron.New(cron.WithLocation(time.UTC))

	_, err := job.AddFunc("30 13 * * *", func() {
		event := events.CreateGameNightEvent(bot)
		events.TransformGameNightToExternalEvent(bot, event)
	})

	if err != nil {
		log.Fatal("Error scheduling cron job", err)
	}

	return job
}
