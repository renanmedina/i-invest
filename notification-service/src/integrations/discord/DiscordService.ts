import DiscordApiClient from "./DiscordApiClient";

const BOT_TOKEN = process.env.DISCORD_BOT_TOKEN || "";
const CHANNEL_ID = process.env.DISCORD_CHANNEL_ID || "";

export default class DiscordService {
  private _api: DiscordApiClient;

  static build() {
    return new DiscordService(
      DiscordApiClient.build(BOT_TOKEN)
    )
  }

  constructor(api: DiscordApiClient) {
    this._api = api;
  }

  async sendMessageToInvestChannel(messageContent: string) {
    const buildedMessage = { content: messageContent, tts: false };
    return await this._api.sendMessage(CHANNEL_ID, buildedMessage);
  }
}