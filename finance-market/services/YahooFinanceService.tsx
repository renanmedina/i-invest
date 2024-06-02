import axios from "axios";
import { IHistoricalDataPrice } from "~/@types/IHistoricalDataPrice";
import { QuoteProps } from "~/@types/QuoteProps";

export default class YahooFinanceService {
  private _http: any;
  private _stockUrl: string = 'https://query1.finance.yahoo.com/v7/finance/options';
  private _historyUrl: string = 'https://query1.finance.yahoo.com/v8/finance/chart';

  constructor(http = axios) {
    this._http = http;
  }

  static build() {
    return new YahooFinanceService();
  }

  async getStockDetailsByCode(code: string): Promise<QuoteProps> {
    const responseYahoo = await this._http.get(`${this._stockUrl}/${code.toUpperCase()}.SA`);
    const result = responseYahoo.data.optionChain.result[0];
    return result.quote as QuoteProps;
  }

  async getStockHistoryData(code: string, interval: string | string[] = '1d', range: string | string[] = '1mo'): Promise<IHistoricalDataPrice[]> {
    try {
      const queryParams = `includePrePost=false&interval=${interval}&useYfid=true&range=${range}`;
      const historicalResponse = await this._http.get(`${this._historyUrl}/${code.toUpperCase()}.SA?${queryParams}`);
      const result = historicalResponse.data.chart.result[0];
      const { timestamp } = result;
      const {
        low,
        high,
        open,
        close,
        volume,
      } = result.indicators.quote[0];

      const { adjclose: adjustedClose } = result.indicators.adjclose[0] || {};

      const prices: Array<IHistoricalDataPrice> = [];

      for (let index = 0; index < timestamp.length; index++) {
        const price: IHistoricalDataPrice = {
          date: timestamp[index],
          open: open[index] || null,
          high: high[index] || null,
          low: low[index] || null,
          close: close[index] || null,
          volume: volume[index] || null,
          adjustedClose: adjustedClose[index] || null,
        };

        prices.push(price);
      }

      return prices;
    } catch (error) {
      console.log(error?.message);
    }
  }
}