import Axios, { AxiosStatic } from 'axios';
import { Logger } from '../../utils/Logger';

const API_ENDPOINT = process.env.DISCORD_API_URL || 'https://discord.com/api/v10';
const CLIENT_ID = process.env.DISCORD_CLIENT_ID || "";
const CLIENT_SECRET = process.env.DISCORD_CLIENT_SECRET || "";
const REDIRECT_URI = process.env.DISCORD_REDIRECT_URI || 'http://localhost:8081/oauth/discord';

export default class DiscordAuthService {
  private _requester: AxiosStatic;
  private _clientId: string;
  private _clientSecret: string;
  private _scopes: string;

  static build() {
    return new DiscordAuthService(
      Axios, 
      CLIENT_ID, 
      CLIENT_SECRET
    );
  }

  constructor(requester: AxiosStatic, clientId: string, clientSecret: string) {
    this._requester = requester;
    this._clientId = clientId;
    this._clientSecret = clientSecret;
    this._scopes = 'identify+bot+messages.read+guilds.join';
  }

  getAuthorizeUrl() {
    const params = [
      `response_type=code`,
      `client_id=${this._clientId}`,
      `scope=${this._scopes}`,
      `permissions=8`,
      `state=i-invest-discord-oauth-state`,
      `redirect_uri=${REDIRECT_URI}`,
    ]

    return `${API_ENDPOINT}/oauth2/authorize?${params.join('&')}`;
  }

  async requestToken(authCode: string) {
    const params = {
      'grant_type': 'authorization_code',
      'code': authCode,
      'redirect_uri': REDIRECT_URI
    }
    return await this._sendTokenRequest(params);
  }

  async refreshToken(refreshToken: string) {
    const params = {
      'grant_type': 'refresh_token',
      'refresh_token': refreshToken
    }
    return await this._sendTokenRequest(params);
  }

  private async _sendTokenRequest(params) {
    const apiUrl = `${API_ENDPOINT}/oauth2/token`;
    Logger.child({apiUrl: apiUrl, params: params}).info("Sending token request to discord service");
    const response = await this._requester.post(
      apiUrl,
      params,
      this._buildConfigs()
    );
    return response.data;
  }

  private _buildConfigs() {
    return { 
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      auth: {
        username: this._clientId,
        password: this._clientSecret
      }
    }
  }
}