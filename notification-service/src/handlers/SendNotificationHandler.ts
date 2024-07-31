import { HttpStatusCode } from "axios";
import { Request, Response } from "express";
import DiscordService from "../integrations/discord/DiscordService";

export default async (req: Request, res: Response) => {
  try {
    const messageToSend = req.query.message || "";
    
    if (messageToSend.length == 0) {
      return res.status(HttpStatusCode.BadRequest).send({
        error: `Message to send is required!`
      });
    }

    const service = DiscordService.build();
    await service.sendMessageToInvestChannel(messageToSend?.toString());
    return res.status(HttpStatusCode.Ok).send({ result: `Notification sent successfully!`});
  } catch (err) {
    return res.status(HttpStatusCode.InternalServerError).send({
      error: `Failed to send notification using notification service`
    })
  }
}