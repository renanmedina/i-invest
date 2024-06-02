import { IHistoricalDataPrice } from '~/@types/IHistoricalDataPrice';
import { QuoteProps } from '~/@types/QuoteProps';

type IValidRangeInterval =
  | '1d'
  | '5d'
  | '1mo'
  | '3mo'
  | '6mo'
  | '1y'
  | '2y'
  | '5y'
  | '10y'
  | 'ytd'
  | 'max';

export interface IQuote extends QuoteProps {
  historicalDataPrice: IHistoricalDataPrice[];
}

interface IGetCurrentQuote {
  stocks: string[] | string;
  range?: IValidRangeInterval;
  interval?: IValidRangeInterval;
  fundamental?: boolean;
}

export const getCurrentQuote = async (props: IGetCurrentQuote) => {
  try {
    const {
      stocks,
      range = '1d',
      interval = '1d',
      fundamental = false,
    } = props;

    const stock = Array.isArray(stocks) ? stocks.join(',') : stocks;

    const url = `https://brapi.dev/api/quote/${stock}?range=${range}&interval=${interval}&fundamental=${fundamental}`;

    const res = await fetch(url, { cache: 'no-cache' });
    const data = (await res.json()) as {
      results: IQuote[];
    };

    return data?.results || [];
  } catch (err) {}
};
