import { NextApiRequest, NextApiResponse } from 'next';
import axios from 'axios';
import { parseDMY } from '~/utils/parseDMY';
import { logHost } from '~/utils/logHost';

const seriesCode = {
  'prime-rate': 432,
  inflation: 13522,
};

const availableSorts = ['date', 'value'];

const availablesortOrder = ['asc', 'desc'];

interface IQuery {
  country?: string;
  historical?: string;
  start?: string;
  end?: string;
  sortBy?: 'date' | 'value';
  sortOrder?: 'asc' | 'desc';
}

interface IResponse {
  conteudo: Array<{
    data: string;
    valor: number;
  }>;
}

export const handleInflationOrPrimeRate = async (
  type: 'inflation' | 'prime-rate',
  req: NextApiRequest,
  res: NextApiResponse,
) => {
  logHost(req, `v2/${type}`);

  let {
    historical = 'false',
    start,
    end,
    sortBy = 'date',
    sortOrder = 'desc',
  } = req.query as IQuery;

  if (sortBy && !availableSorts.includes(sortBy?.toString())) {
    res.status(417).json({
      error: true,
      message: `${sortBy} query value is not available, please use one of the following: ${availableSorts}`,
    });
  }

  if (sortOrder && !availablesortOrder.includes(sortOrder?.toString())) {
    res.status(417).json({
      error: true,
      message: `${sortOrder} query value is not available, please use one of the following: ${availablesortOrder}`,
    });
  }

  if (start || end) {
    historical = 'true';
  }

  if (req.method !== 'GET') {
    res.status(405).send({
      error: true,
      message: 'Method not allowed',
    });
  }

  res.setHeader('Cache-Control', 's-maxage=900, stale-while-revalidate'); // 15 minutes cache

  const datePlusOneMonth = new Date().setMonth(new Date().getMonth() + 1);
  const formattedDatePlusOneMonth = new Date(
    datePlusOneMonth,
  ).toLocaleDateString('pt-BR', {
    timeZone: 'America/Sao_Paulo',
  });
  const startDate =
    historical === 'true' ? '01/01/1500' : formattedDatePlusOneMonth;
  const endDate = formattedDatePlusOneMonth;

  const apiStartDate = start || startDate;
  const apiEndDate = end || endDate;

  try {
    const { data } = await axios.get<IResponse>(
      `https://www.bcb.gov.br/api/servico/sitebcb/bcdatasgs?serie=${seriesCode[type]}&dataInicial=${apiStartDate}&dataFinal=${apiEndDate}`,
    );

    const formattedData = data?.conteudo.map((item) => ({
      date: item.data,
      value: item.valor,
      epochDate: new Date(parseDMY(item.data)).getTime(),
    }));

    const orderedInflationOrPrimeRate = formattedData.sort((a, b) => {
      if (sortBy === 'value') {
        if (sortOrder === 'desc') {
          return b.value - a.value;
        }
        return a.value - b.value;
      }

      if (sortOrder === 'asc') {
        return a.epochDate - b.epochDate;
      }

      return b.epochDate - a.epochDate;
    });

    res.status(200).json({ [type]: orderedInflationOrPrimeRate });
  } catch (err) {
    console.log(err);
    res.status(400).json({
      error: true,
      message: 'Something went wrong while fetching the data',
    });
  }
};
