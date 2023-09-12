# Rambo bot
My Discord bot, its name is Rambo

## Requirements
In able to compile and run project, make sure `python 3.10` or above is installed.

Clone repository by `git clone` to your local computer:
```shell script
$ git clone git@github.com:khoaji/rambo-bot.git
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
1. Create a new virtual environment and activate it
```shell script
$ python -m venv .venv  # python3 in most cases
$ source .venv/bin/activate
```

2. Install all dependencies, this only run once per virtualenv
```shell script
$ pip install -r requirements.txt
```
You might want to run `pip sync` from now on.

3. Compile and run your bot locally:
```shell script
$ python main.py
```

NOTE: Everytime you add a new direct dependency you must re-generate `requirements.txt` by executing `pip-compile` to avoid deployment failure.

## Release
Heroku

## References
[discord.py](https://discordpy.readthedocs.io/en/stable/)
