import axios from 'axios';
import { NextApiRequest, NextApiResponse } from 'next';

interface ILooseObject {
  [key: string]: any;
}

export default async (req: NextApiRequest, res: NextApiResponse) => {
  const data = req.body;
  const formType = req.query.type;

  const formData = JSON.parse(data?.form_data || '{}') as ILooseObject;

  const cleanFormData = Object.keys(formData).reduce((acc, key) => {
    // remove keys that start with _
    if (key.startsWith('_')) {
      return acc;
    }

    return {
      ...acc,
      [key]: formData[key as keyof typeof formData],
    };
  }, {}) as ILooseObject;

  try {
    await axios.post(`${process.env.DISCORD_WEBHOOK_URL}`, {
      username: 'brapi',
      avatar_url: 'https://brapi.dev/favicon.png',
      embeds: [
        {
          title: `New ${formType} Form Submission`,
          description: cleanFormData?.email || '',
          color: 7419530,
          fields: Object.entries(cleanFormData || {}).map(([key, value]) => ({
            name: key,
            value,
          })),
        },
      ],
    });

    console.log('webhook sent to discord');

    res.status(200).json({ message: 'ok' });
  } catch (error) {
    console.error({ error });
    res.status(500).json({ error: true });
  }
};
