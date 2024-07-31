import { Request, Response } from "express";
import { HttpStatusCode } from "axios";
import DiscordAuthService from "../../integrations/discord/DiscordAuthService";

export default async (req: Request, res: Response) => {
  const authCode = req.query.code;
  const authState = req.query.state;
  const discordService = DiscordAuthService.build();

  if (authCode && authState) {
    try {
      const tokenData = await discordService.requestToken(authCode.toString());
      return res.status(HttpStatusCode.Ok).send({ result: tokenData });
    } catch (exception) {
      return res.status(HttpStatusCode.InternalServerError).send({
        error: 'Failed to exchange authorization code for access token on discord service'
      });
    }
  }

  return res.redirect(discordService.getAuthorizeUrl());
}
