import { Client } from "discord.js";
import { Commads } from "src/commands/commands";
import { logger } from "src/utils/logger";

export default (client: Client): void => {
  client.once("ready", async () => {
    if (!client.user || !client.application) {
      return;
    }

    await client.application.commands.set(Commads);

    logger.info(`Ready! Logged in as ${client.user.tag}`);
  });
};
