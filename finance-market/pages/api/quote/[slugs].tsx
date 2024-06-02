import { logHost } from '../../../utils/logHost';
import { NextApiRequest, NextApiResponse } from 'next';
import { QuoteProps } from '../../../@types/QuoteProps';
import { TradingViewQuoteResponse } from '~/@types/TradingViewQuoteResponse';
import TradingViewService from '~/services/TradingViewService';
import DividendsService from '~/services/DividendsService';
import YahooFinanceService from '~/services/YahooFinanceService';

interface LooseObject {
  [key: string]: any;
}

export default async (req: NextApiRequest, res: NextApiResponse) => {
  logHost(req, 'quote');

  const { slugs, interval, range, fundamental, dividends } = req.query;

  const validRanges = [
    '1d',
    '5d',
    '1mo',
    '3mo',
    '6mo',
    '1y',
    '2y',
    '5y',
    '10y',
    'ytd',
    'max',
  ];

  const allSlugs = slugs.toString().split(',');
  const financeService = YahooFinanceService.build(); 

  if (slugs) {
    const responseAllSlugs = async () => {
      const promises = allSlugs.map(async (slug) => {
        try {
          const data: QuoteProps = await financeService.getStockDetailsByCode(slug);
          let quote: LooseObject = {
            symbol: slug.toString().toUpperCase(),
            shortName: data.shortName,
            longName: data.longName,
            currency: data.currency,
            regularMarketPrice: data.regularMarketPrice,
            regularMarketDayHigh: data.regularMarketDayHigh,
            regularMarketDayLow: data.regularMarketDayLow,
            regularMarketDayRange: data.regularMarketDayRange,
            regularMarketChange: data.regularMarketChange,
            regularMarketChangePercent: data.regularMarketChangePercent,
            regularMarketTime: new Date(data.regularMarketTime * 1000),
            marketCap: data.marketCap,
            regularMarketVolume: data.regularMarketVolume,
            regularMarketPreviousClose: data.regularMarketPreviousClose,
            regularMarketOpen: data.regularMarketOpen,
            averageDailyVolume10Day: data.averageDailyVolume10Day,
            averageDailyVolume3Month: data.averageDailyVolume3Month,
            fiftyTwoWeekLowChange: data.fiftyTwoWeekLowChange,
            fiftyTwoWeekLowChangePercent: data.fiftyTwoWeekLowChangePercent,
            fiftyTwoWeekRange: data.fiftyTwoWeekRange,
            fiftyTwoWeekHighChange: data.fiftyTwoWeekHighChange,
            fiftyTwoWeekHighChangePercent: data.fiftyTwoWeekHighChangePercent,
            fiftyTwoWeekLow: data.fiftyTwoWeekLow,
            fiftyTwoWeekHigh: data.fiftyTwoWeekHigh,
            twoHundredDayAverage: data.twoHundredDayAverage,
            twoHundredDayAverageChange: data.twoHundredDayAverageChange,
            twoHundredDayAverageChangePercent:
              data.twoHundredDayAverageChangePercent,
          };

          if (fundamental) {
            try {
              const fundamentalInformation: TradingViewQuoteResponse = await TradingViewService.build().getByTickerCode(slug);
              quote.priceEarnings = fundamentalInformation.price_earnings_ttm;
              quote.earningsPerShare = fundamentalInformation.earnings_per_share_basic_ttm;
              quote.type = fundamentalInformation.type;
              quote.sector = fundamentalInformation.sector;
              quote.logourl = fundamentalInformation.logoid
                ? `https://s3-symbol-logo.tradingview.com/${fundamentalInformation.logoid}--big.svg`
                : 'https://brapi.dev/favicon.svg';
            } catch (error) {
              console.log(error?.message);
            }
          }

          if (dividends) {
            try {
              quote.dividendsData = await DividendsService.build().getByTickerCode(slug);
            } catch (error) {
              console.log(error?.message);
            }
          }

          if (interval && range) {
            const historicalData = await financeService.getStockHistoryData(slug, interval, range);
            quote.historicalDataPrice = historicalData;
            quote.validRanges = validRanges;
          }

          return quote;
        } catch (err) {
          return {
            symbol: slug.toString().toUpperCase(),
            error: true,
            message: `Não encontramos a ação ${slug.toString().toUpperCase()}`,
          };
        }
      });

      const dynamicDate = new Date();
      await Promise.all(promises)
        .then((actualData) => {
          if (actualData?.length === 1 && actualData?.[0]?.error) {
            throw new Error(actualData[0].message);
          }

          res.setHeader(
            'Cache-Control',
            's-maxage=900, stale-while-revalidate',
          );

          res.status(200).json({
            results: actualData,
            requestedAt: dynamicDate,
          });
        })
        .catch((err) => {
          res.setHeader('Cache-Control', 's-maxage=10, stale-while-revalidate');
          return res.status(404).json({
            error: err.message,
          });
        });
    };

    await responseAllSlugs();
  }
};
