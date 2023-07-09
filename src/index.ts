import { Client, Collection, GatewayIntentBits } from "discord.js";
import * as config from "./config.json";
import ready from "./events/ready";
import interactionCreate from "./events/interaction-create";

const client = new Client({
  intents: [
    GatewayIntentBits.Guilds,
    GatewayIntentBits.GuildMessages,
    GatewayIntentBits.MessageContent,
  ],
});

ready(client);
interactionCreate(client)

// Login to Discord
client.login(config.token);
