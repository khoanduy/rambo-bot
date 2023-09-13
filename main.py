from core import config, bot
from utils.logger import LOGGER

bot = bot.Bot()

if __name__ == '__main__':
    if config.TOKEN is None:
        LOGGER.error('Could not find any TOKEN environment variable')
        exit()

    bot.run(config.TOKEN)
