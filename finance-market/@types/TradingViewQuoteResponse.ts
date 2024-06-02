export interface TradingViewQuoteResponse {
  stock_code: string,
  close?: number,
  change?: number,
  volume?: number,
  market_cap_basic?: number,
  description: string, 
  logoid: string,
  type: string,
  sector: string,
  price_earnings_ttm: number,
  earnings_per_share_basic_ttm: number
}

