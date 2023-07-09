package slash

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "spam",
		Description: "Spam a specific user",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "User to spam",
				Required:    true,
			},
		},
	},
}
