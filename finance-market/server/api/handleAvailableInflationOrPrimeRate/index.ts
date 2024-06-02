import { NextApiRequest, NextApiResponse } from 'next';
import { availableInflationCountries } from '~/utils/availableInflationCountries';
import { logHost } from '~/utils/logHost';

export const handleAvailableInflationOrPrimeRate = async (
  type: 'inflation' | 'prime-rate',
  req: NextApiRequest,
  res: NextApiResponse,
) => {
  logHost(req, `v2/${type}/available`);

  const { search } = req.query;

  res.setHeader('Cache-Control', 's-maxage=2592000, stale-while-revalidate');

  if (!!search?.length) {
    const filteredCountries = availableInflationCountries.filter((key) =>
      key.toLowerCase().includes(search.toString().toLowerCase()),
    );

    if (!!filteredCountries?.length) {
      res.status(200).json({
        countries: filteredCountries,
      });
      return;
    }

    res.status(404).json({
      message: 'Country not found',
    });
    return;
  }

  res.status(200).json({
    countries: availableInflationCountries,
  });
};
