const { Events } = require('discord.js');
const logger = require('../utils/logger');
const { guildId, voiceChannelId, voiceChannelLink } = require('../config.json');

module.exports = {
  name: Events.ClientReady,
  once: true,
  execute(client) {
    logger.info(`Ready! Logged in as ${client.user.tag}`);

    // Create event at 20:30 UTC everyday
    const guild = client.guilds.cache.get(guildId);
    const voiceChannel = guild.channels.cache.get(voiceChannelId);

    console.log(guild.scheduledEvents);

    guild.scheduledEvents.create({
      name: 'Game Night Event',
      description: 'This event will start in 30 minutes and last for 3 hours.',
      scheduledStartTime: new Date(Date.now() + 30 * 60000),
      scheduledEndTime: new Date(Date.now() + 3 * 3600000),
      entityType: 'voice',
      entityMetadata: {
        channelId: voiceChannelId,
        privacyLevel: 'guild_only',
        location: voiceChannelLink,
      },
    }).then(scheduledEvent => {
      logger.info('Created scheduled event:', scheduledEvent.name);
    }).catch(error => {
      logger.error('Error creating scheduled event:', error);
    });
  },
};
