package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/samothrakii/rambo-bot/conf"
	"github.com/samothrakii/rambo-bot/core"
)

func main() {
	conf.LoadConf()

	bot := core.BotInit()
	bot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	core.RegisterSlashCommands(bot)
	core.RegisterMessageReplyBehavior(bot)

	err := bot.Open()
	if err != nil {
		fmt.Println("Error opening connection", err)
		return
	}

	gameNightEventCron := core.GameNightEventCron(bot)
	gameNightEventCron.Start()

	fmt.Println("Rambo is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	gameNightEventCron.Stop()
	bot.Close()

	log.Println("Gracefully shutting down.")
	os.Exit(0)
}
