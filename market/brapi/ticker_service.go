package brapi

import (
	"fmt"
)

type TickerService struct {
	client ApiClient[Ticker]
}

type PriceHistory struct {
}

type Ticker struct {
	Code           string         `json:"symbol"`
	LogoUrl        string         `json:"logourl"`
	Name           string         `json:"shortName"`
	Currency       string         `json:"currency"`
	Price          float64        `json:"regularMarketPrice"`
	LastClosePrice float64        `json:"regularMarketPreviousClose"`
	PriceHistory   []PriceHistory `json:"historicalDataPrice"`
}

func NewTicketService() TickerService {
	return TickerService{client: NewApiClient[Ticker]()}
}

func (ts *TickerService) GetByCode(tickerCode string) (Ticker, error) {
	var tickers []Ticker
	response := ts.client.Get(fmt.Sprintf("/quote/%s?fundamental=true&dividends=true", tickerCode))
	tickers = response.Results
	if len(tickers) == 0 {
		return Ticker{}, fmt.Errorf("Ticker %s não encontrado", tickerCode)
	}

	return tickers[0], nil
}
