package events

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/khoaji/rambo-bot/config"
)

func CreateGameNightEvent(s *discordgo.Session) *discordgo.GuildScheduledEvent {
	startingTime := time.Now().Add(30 * time.Minute)
	endingTime := startingTime.Add(3 * time.Hour)
	scheduledEvent, err := s.GuildScheduledEventCreate(config.Env.GuildId, &discordgo.GuildScheduledEventParams{
		Name:               "Game Night Event",
		Description:        "This event will start in 30 mins and last 3 hours",
		ScheduledStartTime: &startingTime,
		ScheduledEndTime:   &endingTime,
		EntityType:         discordgo.GuildScheduledEventEntityTypeVoice,
		ChannelID:          config.Env.VoiceChanId,
		PrivacyLevel:       discordgo.GuildScheduledEventPrivacyLevelGuildOnly,
	})

	if err != nil {
		log.Printf("Error creating scheduled event: %v", err)
		return nil
	}

	s.ChannelMessageSend(config.Env.ChannelId, "Game Night event starts in 30 mins. Get ready @here.")

	fmt.Println("Created scheduled event:", scheduledEvent.Name)
	return scheduledEvent
}

func TransformGameNightToExternalEvent(s *discordgo.Session, event *discordgo.GuildScheduledEvent) {
	scheduledEvent, err := s.GuildScheduledEventEdit(config.Env.GuildId, event.ID, &discordgo.GuildScheduledEventParams{
		Name:       "Game Night @ Hamster Hill",
		EntityType: discordgo.GuildScheduledEventEntityTypeExternal,
		EntityMetadata: &discordgo.GuildScheduledEventEntityMetadata{
			Location: config.Env.VoiceChanLink,
		},
	})
	if err != nil {
		log.Printf("Error during transformation of scheduled voice event into external event: %v", err)
		return
	}

	fmt.Println("Transformed scheduled event:", scheduledEvent.Name)
}
