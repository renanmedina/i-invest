import { NextApiRequest, NextApiResponse } from 'next';
import { handleAvailableInflationOrPrimeRate } from '~/server/api/handleAvailableInflationOrPrimeRate';

export default async (req: NextApiRequest, res: NextApiResponse) => {
  return handleAvailableInflationOrPrimeRate('inflation', req, res);
};
