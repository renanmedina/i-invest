import express from 'express';
import 'dotenv/config';
import { HttpLogger, Logger } from './utils/Logger';
import SendNotificationHandler from './handlers/SendNotificationHandler';
import OAuthDiscordHandler from './handlers/discord/OAuthDiscordHandler';

const server = express();
const PORT = process.env.PORT || 8081

server.use(HttpLogger);
server.get('/notify', SendNotificationHandler);
server.get('/oauth/discord', OAuthDiscordHandler);

server.listen(PORT, () => {
  Logger.info(`Notification service webserver running on port ${PORT}`);
});