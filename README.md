# rambo-bot
My Discord bot, its name is Rambo

## Requirements
In able to run project, make sure `python 3.10` or above is installed.

Clone repository by `git clone` to your local computer:
```shell script
$ git clone git@github.com:khoanduy/rambo-bot.git
$ cd rambo-bot
```

You need to add these following values to `.env` file
```
TOKEN="your_bot_token"
APP_ID="your_bot_application_id"
GUILD_ID="your_target_guild_id"
CHANNEL_ID="your_target_channel_id"
DEV_CHANNEL_ID="your_target_development_channel_id"
```

## Build and run
1. Create new virtual environment and activate it:
```shell script
$ python -m venv venv
$ source venv/bin/activate
```
2. For development install all dependencies by executing:
```shell script
$ pip install -r dev-requirements.txt
```
3. Launch the bot:
```shell script
$ python main.py
```

## Release
TBA

## Additional notes
This project use `pip-tools` to manage dependencies, to install simply run the following command in your venv:
```shell script
$ pip install pip-tools
```

To update dependencies, just add them into `[project.dependencies]` section in `pyproject.toml` and run:
```shell script
$ pip-compile -o requirements.txt pyproject.toml
```

For development dependencies, update `[project.optional-dependencies.dev]` section instead and run:
```shell script
$ pip-compile --extra dev -o dev-requirements.txt pyproject.toml
```

Make sure those files are committed with your changes.

## References
[Discord docs](https://discord.com/developers/docs)
