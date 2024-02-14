# rambo-bot
My Discord bot, its name is Rambo

## Requirements
In able to run project, make sure `rustc 1.76.0` or above is installed.

Clone repository by `git clone` to your local computer:
```shell script
$ git clone git@github.com:khoanduy/rambo-bot.git
$ cd rambo-bot
```

You need to add these following values to `.cargo/config.toml` file
```toml
[env]
TOKEN="your_bot_token"
APP_ID="your_bot_application_id"
GUILD_ID="your_target_guild_id"
CHANNEL_ID="your_target_channel_id"
DEV_CHANNEL_ID="your_target_development_channel_id"
```

## Build and run
Install all dependencies and build project
```shell script
$ cargo build
```
Run your bot locally
```shell script
$ cargo run
```

## Release
TBA

## References
[serenity](https://github.com/serenity-rs/serenity)
