import pino from 'pino';
import pinoHttp from 'pino-http';

export const Logger = pino();
export const HttpLogger = pinoHttp();