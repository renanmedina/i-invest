package brapi

import (
	"fmt"
	"strings"
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

func NewTickerService() TickerService {
	return TickerService{client: NewApiClient[Ticker]()}
}

func (ts *TickerService) GetByCode(tickerCode string) (Ticker, error) {
	codes := []string{tickerCode}
	tickers, err := ts.GetByCodes(codes)

	if err != nil {
		return Ticker{}, fmt.Errorf("ticker %s não encontrado", tickerCode)
	}

	return tickers[0], nil
}

func (ts *TickerService) GetByCodes(tickerCodes []string) ([]Ticker, error) {
	var tickers []Ticker
	codes := strings.Join(tickerCodes, ",")
	response := ts.client.Get(fmt.Sprintf("/quote/%s?fundamental=true&dividends=true", codes))
	tickers = response.Results

	if len(tickers) == 0 {
		return tickers, fmt.Errorf("tickers %s não encontrados", codes)
	}

	return tickers, nil
}
