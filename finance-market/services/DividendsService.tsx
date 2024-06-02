import axios from "axios";
import { Dividends } from "~/@types/Dividends";
import { parseDMY } from '~/utils/parseDMY';
import { replaceComma } from '~/utils/replaceComma';

export default class DividendsService {
  private _http: any;
  private _baseUrl: string = 'https://sistemaswebb3-listados.b3.com.br/fundsProxy/fundsCall/GetListedSupplementFunds';

  constructor(http = axios) {
    this._http = http;
  }

  static build() {
    return new DividendsService();
  }

  async getByTickerCode(code: string): Promise<Dividends> {
    const jwtHeaderString = this._makeJwt(code);
    const responseDividends = await this._http.get(`${this._baseUrl}/${jwtHeaderString}`);
    const { cashDividends, stockDividends, subscriptions } = responseDividends?.data || {};

    return {
      cashDividends: cashDividends?.map(this.dividendParser),
      stockDividends: stockDividends?.map(this.dividendParser),
      subscriptions: subscriptions?.map(this.dividendParser),
    } as Dividends
  }

  dividendParser(eachDividend) {
    return {
      ...eachDividend,
      ...(eachDividend?.paymentDate && {
        paymentDate: new Date(parseDMY(eachDividend?.paymentDate)),
      }),
      ...(eachDividend?.approvedOn && {
        approvedOn: new Date(parseDMY(eachDividend?.approvedOn)),
      }),
      ...(eachDividend?.lastDatePrior && {
        lastDatePrior: new Date(
          parseDMY(eachDividend?.lastDatePrior),
        ),
      }),
      ...(eachDividend?.rate && {
        rate: parseFloat(replaceComma(eachDividend?.rate)),
      }),
      ...(eachDividend?.factor && {
        factor: parseFloat(replaceComma(eachDividend?.factor)),
      }),
      ...(eachDividend?.percentage && {
        percentage: parseFloat(
          replaceComma(eachDividend?.percentage),
        ),
      }),
      ...(eachDividend?.priceUnit && {
        priceUnit: parseFloat(
          replaceComma(eachDividend?.priceUnit),
        ),
      }),
      ...(eachDividend?.subscriptionDate && {
        subscriptionDate: new Date(
          parseDMY(eachDividend?.subscriptionDate),
        ),
      }),
    };
  };

  _makeJwt(tickerCode: string) {
    const jwtHeader = {
      identifierFund: tickerCode,
    };

    return Buffer.from(JSON.stringify(jwtHeader)).toString('base64');
  }
}