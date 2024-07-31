import Axios, { AxiosStatic } from 'axios';
import DiscordMessageBody from './interfaces/DiscordMessageBody';

const API_ENDPOINT = process.env.DISCORD_API_URL || 'https://discord.com/api/v10';

export default class DiscordApiClient {
  private _requester: AxiosStatic;
  private _accessToken: string;

  static build(accessToken: string) {
    return new DiscordApiClient(Axios, accessToken)
  }

  constructor(requester: AxiosStatic, accessToken: string) {
    this._requester = requester;
    this._accessToken = accessToken;
  }

  async sendMessage(channelId: string, messageBody: DiscordMessageBody) {
    return await this._requester.post(
      `${API_ENDPOINT}/channels/${channelId}/messages`,
      messageBody,
      this._buildHeaders()
    )
  }

  private _buildHeaders() {
    return {
      headers: {
        'Authorization': this._accessToken,
        'Content-Type': 'application/json'
      }
    }
  }
}