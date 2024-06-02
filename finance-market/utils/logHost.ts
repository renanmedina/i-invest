import axios from 'axios';
import { getClientIp } from 'request-ip';
import { NextApiRequest } from 'next';

const measurementId = process.env.NEXT_PUBLIC_GTAG_ID;
const apiSecret = process.env.ANALYTICS_SECRET_KEY;
const clientId = process.env.ANALYTICS_CLIENT_ID;
const url = `https://www.google-analytics.com/mp/collect?measurement_id=${measurementId}&api_secret=${apiSecret}`;

export const logHost = async (req: NextApiRequest, path: string) => {
  const { referer, origin } = req.headers;
  const userIp = getClientIp(req);

  const body = JSON.stringify({
    client_id: clientId,
    events: [
      {
        name: 'api_usage',
        params: {
          endpoint: path,
          origin: origin || referer || userIp,
        },
      },
      {
        name: `${path}_endpoint`,
        params: {
          fullpath: req.url?.replace('api/', ''),
          origin: origin || referer || userIp,
        },
      },
    ],
  });

  try {
    await axios.post(url, body);
  } catch (e) {
    return;
  }
};
