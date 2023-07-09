package core

import (
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/khoaji/rambo-bot/utils"
)

func MessageReply(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if isMsgUrl(m.Content) {
		links := findUrlsInMsg(m.Content)

		var safe []string
		var unsafe []string
		for _, link := range links {
			if isSafe, err := utils.IsSafeLink(link); err != nil {
				log.Println("Could not check link", err)
				return
			} else if isSafe {
				safe = append(safe, link)
			} else {
				unsafe = append(unsafe, link)
			}
		}

		res := "Link(s) "
		if len(safe) > 0 {
			res += "safe to click\n"
		}

		if len(unsafe) > 0 {
			res += strings.Join(unsafe, ", ")
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
