import axios from 'axios';
import { NextApiRequest, NextApiResponse } from 'next';
import { logHost } from '../../../../utils/logHost';

interface DataProps {
  [key: string]: string;
}

export default async (req: NextApiRequest, res: NextApiResponse) => {
  logHost(req, 'v2/crypto/available');

  const { search } = req.query;

  res.setHeader('Cache-Control', 's-maxage=2592000, stale-while-revalidate');

  const { data } = await axios.get<DataProps>(
    'https://economia.awesomeapi.com.br/json/available',
  );

  if (!!search?.length) {
    const filteredCurrencies = Object.keys(data).filter(
      (key) =>
        key.toLowerCase().includes(search.toString().toLowerCase()) ||
        data[key].toLowerCase().includes(search.toString().toLowerCase()),
    );

    const currencies = filteredCurrencies.map((key) => ({
      name: key,
      currency: data[key],
    }));

    if (!!currencies?.length) {
      res.status(200).json({
        currencies,
      });
      return;
    }

    res.status(404).json({
      message: 'Currency not found',
    });
    return;
  }

  const currenciesData = Object.keys(data).map((key) => ({
    name: key,
    currency: data[key],
  }));

  res.status(200).json({
    currencies: currenciesData,
  });
};
