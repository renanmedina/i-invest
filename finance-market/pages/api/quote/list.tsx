import { logHost } from '../../../utils/logHost';
import { NextApiRequest, NextApiResponse } from 'next';
import TradingViewService from '~/services/TradingViewService';
import { TradingViewQuoteResponse } from '~/@types/TradingViewQuoteResponse';

export default async (req: NextApiRequest, res: NextApiResponse) => {
  logHost(req, 'list');

  const { sortBy, sortOrder, limit, search } = req.query;

  res.setHeader('Cache-Control', 's-maxage=900, stale-while-revalidate');

  const availableSorts = [
    'name',
    'close',
    'change',
    'change_abs',
    'volume',
    'market_cap_basic',
  ];

  const availablesortOrder = ['desc', 'asc'];

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

  if (limit && !Number(limit)) {
    res.status(417).json({
      error: true,
      message: `limit must be a number`,
    });
  }

  const stockItems = await TradingViewService.build().getList({
    sortBy: sortBy,
    sortOrder: sortOrder,
    search: search,
    limit: limit
  });

  const paths = stockItems.map((item: TradingViewQuoteResponse) => {
    const logo = item.logoid ? `https://s3-symbol-logo.tradingview.com/${item.logoid}--big.svg` : 'https://brapi.dev/favicon.svg'

    return {
      stock: item.stock_code,
      name: item.description,
      type: item.type,
      close: item.close,
      change: item.change,
      volume: item.volume,
      market_cap: item.market_cap_basic,
      logo: logo,
      sector: item.sector,
    }
  });

  const uniqueStocks = [...new Set(paths)];

  res.status(200).json({
    stocks: uniqueStocks,
  });
};
