package brapi

import (
	"fmt"
	"strings"

	"github.com/renanmedina/investment-warlock/internal/market"
)

type TickerService struct {
	client market.ApiClient[TickerApiResult]
}

type PriceHistory struct {
}

type TickerApiResult struct {
	Results []Ticker `json:"results"`
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
	client := market.NewApiClient[TickerApiResult](market.ApiConfig{
		ApiUrl: "https://brapi.dev/api",
	})
	return TickerService{client}
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
	response, err := ts.client.Get(fmt.Sprintf("/quote/%s?fundamental=true&dividends=true", codes), make(map[string]string))

	if err != nil {
		return make([]Ticker, 0), err
	}

	tickers = response.Results

	if len(tickers) == 0 {
		return tickers, fmt.Errorf("tickers %s não encontrados", codes)
	}

	return tickers, nil
}

func (ts *TickerService) GetPricesByCodes(tickerCodes []string) (map[string]float64, error) {
	prices := make(map[string]float64)

	tickers, err := ts.GetByCodes(tickerCodes)

	if err != nil {
		return prices, err
	}

	for _, ticker := range tickers {
		prices[ticker.Code] = ticker.Price
	}

	return prices, nil
}
