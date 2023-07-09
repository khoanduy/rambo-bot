package core

import (
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/samothrakii/rambo-bot/utils"
)

func MessageReply(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if isMsgUrl(m.Content) {
		links := findUrlsInMsg(m.Content)

		var safeLinks []string
		var unsafeLinks []string
		for _, link := range links {
			unsafe, err := utils.CheckUnsafeLink(link)
			if err != nil {
				log.Println("Could not check link ", err)
				return
			}

			if unsafe {
				unsafeLinks = append(unsafeLinks, link)
			} else {
				safeLinks = append(safeLinks, link)
			}
		}

		res := "Link(s) "
		if len(safeLinks) > 0 {
			res += "safe to click\n"
		}

		if len(unsafeLinks) > 0 {
			res += strings.Join(unsafeLinks, ", ")
			res += " NOT SAFE, watch out @here"
		}

		s.ChannelMessageSend(m.ChannelID, res)
	}
}

func isMsgUrl(msg string) bool {
	return regexp.MustCompile(`https?://`).MatchString(msg)
}

func findUrlsInMsg(msg string) []string {
	return regexp.MustCompile(`(?i)\bhttps?://\S+\b`).FindAllString(msg, -1)
}
