package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/khoaji/rambo-bot/config"
	"github.com/khoaji/rambo-bot/core"
)

func main() {
	config.LoadConf()

	bot := core.BotInit()
	bot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	core.RegisterSlashCommands(bot)
	core.RegisterMessageReplyBehavior(bot)

	err := bot.Open()
	if err != nil {
		log.Fatal("Error opening connection", err)
		return
	}

	gameNightEventCron := core.GameNightEventCron(bot)
	gameNightEventCron.Start()

	log.Println("Rambo is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	gameNightEventCron.Stop()
	bot.Close()

	log.Println("Gracefully shutting down.")
	os.Exit(0)
}
