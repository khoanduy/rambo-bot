import discord
from core import config
from core.logger import LOGGER

intents = discord.Intents.default()
intents.message_content = True

client = discord.Client(intents=intents)


@client.event
async def on_ready():
    LOGGER.info(f'We have logged in as {client.user}')


@client.event
async def on_message(message):
    if message.author == client.user:
        return

    if message.content.startswith('$hello'):
        await message.channel.send('Hello!')


if __name__ == '__main__':
    if config.TOKEN is None:
        LOGGER.error('Could not find any TOKEN environment variable')
        exit()

    client.run(config.TOKEN)
