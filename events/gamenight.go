package events

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/khoaji/rambo-bot/config"
)

func CreateGameNightEvent(s *discordgo.Session) *discordgo.GuildScheduledEvent {
	start := time.Now().Add(30 * time.Minute)
	end := start.Add(3 * time.Hour)
	scheduledEvent, err := s.GuildScheduledEventCreate(config.Env.GuildId, &discordgo.GuildScheduledEventParams{
		Name:               "Game Night Event",
		Description:        "This event will start in 30 mins and last 3 hours",
		ScheduledStartTime: &start,
		ScheduledEndTime:   &end,
		EntityType:         discordgo.GuildScheduledEventEntityTypeVoice,
		ChannelID:          config.Env.VoiceChannelId,
		PrivacyLevel:       discordgo.GuildScheduledEventPrivacyLevelGuildOnly,
	})

	if err != nil {
		log.Printf("Error creating scheduled event: %v", err)
		return nil
	}

	s.ChannelMessageSend(config.Env.ChannelId, "Game Night event starts in 30 mins. Get ready @here.")

	log.Println("Created scheduled event: ", scheduledEvent.Name)
	return scheduledEvent
}

func TransformGameNightToExternalEvent(s *discordgo.Session, event *discordgo.GuildScheduledEvent) {
	scheduledEvent, err := s.GuildScheduledEventEdit(config.Env.GuildId, event.ID, &discordgo.GuildScheduledEventParams{
		Name:       "Game Night @ Hamster Hill",
		EntityType: discordgo.GuildScheduledEventEntityTypeExternal,
		EntityMetadata: &discordgo.GuildScheduledEventEntityMetadata{
			Location: config.Env.VoiceChannelLink,
		},
	})

	if err != nil {
		log.Printf("Error during transformation of scheduled voice event into external event: %v", err)
		return
	}

	log.Println("Transformed scheduled event: ", scheduledEvent.Name)
}
