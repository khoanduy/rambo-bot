import { ApplicationCommandType, Client, CommandInteraction, SlashCommandBuilder } from "discord.js";
import { Command } from "src/interfaces/command";

// module.exports = {
//   data: new SlashCommandBuilder()
//     .setName('spam')
//     .setDescription('Spam a specific user')
//     .addUserOption(option =>
//       option.setName('user')
//         .setDescription('User to spam')
//         .setRequired(true)),
//   async execute(interaction) {
//     const user = interaction.options.get('user').user;
//     await interaction.reply(`${user} `.repeat(20));
//   },
// };

export const Spam: Command = {
  name: "spam",
  description: "Spam a specific user",
  type: ApplicationCommandType.ChatInput,
  run: async (client: Client, interaction: CommandInteraction) => {
    const content = "Hello there!";

    await interaction.followUp({
      ephemeral: true,
      content
    });
  }
};
