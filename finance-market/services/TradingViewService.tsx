import axios from "axios";
import { TradingViewQuoteResponse } from "~/@types/TradingViewQuoteResponse";

const DEFAULT_TICKER_COLUMNS = [
  'close',
  'change',
  'volume',
  'market_cap_basic',
  'description',
  'price_earnings_ttm',
  'earnings_per_share_basic_ttm',
  'logoid',
  'type',
  'sector'
]

interface TradingViewQuoteItem {
  d: any[]
  s: string
}

export default class TradingViewService {
  private _http: any;
  private _baseUrl = 'https://scanner.tradingview.com/brazil/scan';

  constructor(http = axios) {
    this._http = http;
  }

  static build() {
    return new TradingViewService();
  }

  async getByTickerCode(code: string): Promise<TradingViewQuoteResponse> {
    const responseTradingViewData = await this._postRequest({
      symbols: {
        tickers: [`BMFBOVESPA:${code.toUpperCase()}`],
        query: {
          types: [],
        },
      },
      columns: DEFAULT_TICKER_COLUMNS
    });

    return this._mapResponseObject(DEFAULT_TICKER_COLUMNS, responseTradingViewData.pop());
  }

  async getList({sortBy, sortOrder, limit, search = null}): Promise<TradingViewQuoteResponse[]> {
    const filters = [
      {
        left: sortBy?.toString() || 'volume',
        operation: 'nempty',
        right: '',
      },
    ];

    if (search) {
      filters.push({
        left: 'name',
        operation: 'match',
        right: search.toString(),
      });
    }

    const stocks: TradingViewQuoteItem[] = await this._postRequest({
      filter: filters,
      columns: DEFAULT_TICKER_COLUMNS,
      sort: {
        sortBy: sortBy?.toString() || 'volume',
        sortOrder: sortOrder?.toString() || 'desc',
      },
      range: [0, Number(limit) || 2000],
    });

    return stocks.map((quoteItem) => this._mapResponseObject(DEFAULT_TICKER_COLUMNS, quoteItem));
  }

  async _postRequest({filter = [], symbols = {}, columns = DEFAULT_TICKER_COLUMNS, sort = null, range = [0, 2000]}): Promise<TradingViewQuoteItem[]> {
    const params: any = {
      filter: filter,
      symbols: symbols,
      options: {
        lang: 'pt',
        active_symbols_only: true,
      },
      columns: columns,
      range: range,
    };

    if (sort) {
      params.sort = sort;
    }
    
    const responseTradingView = await this._http.post(
      this._baseUrl,
      params,
      {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
      },
    );

    const responseData = responseTradingView.data;
    return responseData.data;
  }

  _mapResponseObject(columns, responseItem: TradingViewQuoteItem): TradingViewQuoteResponse {
    let responseObject: TradingViewQuoteResponse = {
      stock_code: responseItem.s.replace('BMFBOVESPA:', '')
    } as TradingViewQuoteResponse;

    let i = 0;
    while(i < columns.length) {
      const columnName = columns[i];
      let columnValue = responseItem.d[i];
      if (columnName == "description") {
        columnValue = this._cleanString(columnValue);
      }
      responseObject[columnName] = columnValue;
      i++;
    }
    
    return responseObject;
  }

  _cleanString(dirtyString: string) {
    return dirtyString
      .replace(' ON', '')
      .replace(' ON', '')
      .replace(' NM', '')
      .replace(' EJ', '')
      .replace(' REON', '')
      .replace(' N1', '')
      .replace(' N2', '');
  }
}