import os
from dotenv import load_dotenv

load_dotenv()

TOKEN = os.getenv('TOKEN')
APP_ID = os.getenv('APP_ID')
GUILD_ID = os.getenv('GUILD_ID')
CHANNEL_ID = os.getenv('CHANNEL_ID')
DEV_CHANNEL_ID = os.getenv('DEV_CHANNEL_ID')
VOICE_CHANNEL_ID = os.getenv('VOICE_CHANNEL_ID')
VOICE_CHANNEL_LINK = os.getenv('VOICE_CHANNEL_LINK')
VT_API_KEY = os.getenv('VT_API_KEY')
QUIZ_EMBED_IMG_URL = os.getenv('QUIZ_EMBED_IMG_URL')

