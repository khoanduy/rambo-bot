# Rambo bot
My Discord bot, its name is Rambo

## Requirements
In able to compile and run project, make sure `go 1.20` or above is installed.

Clone repository by `git clone` to your local computer:
```shell script
$ git clone git@github.com:khoarc/rambo-bot.git
$ cd rambo-bot
```

You need to add these following environment variables to `.env` (recommended) or `~/.profile` file
```
TOKEN=[bot_token]
MONGO_URI=[mongodb_uri]
APP_ID=[application_id]
GUILD_ID=[target_guild_id]
CHANNEL_ID=[target_channel_id]
DEV_CHANNEL_ID=[dev_channel_id]
VOICE_CHANNEL_ID=[target_voice_channel_id]
VOICE_CHANNEL_LINK=[target_voice_channel_link]
VT_API_KEY=[virus_total_api_key]
```

## Build and run
Install all dependencies:
```shell script
$ go install
```
Compile and run your bot locally:
```shell script
$ go run main.go
```

## Release
To build a single executable binary, simply run:
```shell script
$ go build -ldflags '-s -w' -o bin/rambo-bot
$ ./bin/rambo-bot
```

## References
[discordgo](https://github.com/bwmarrin/discordgo)
