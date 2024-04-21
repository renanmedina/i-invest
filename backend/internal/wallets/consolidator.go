package wallets

import (
	"fmt"
	"math"

	"github.com/renanmedina/investment-warlock/internal/market/brapi"
)

type ConsolidatorItem struct {
	Grouper             string
	TotalQuantity       int
	AveragePrice        float64
	PaidCost            float64
	AverageAmount       float64
	CurrentAmount       float64
	WalletPercentage    float64
	Details             string
	VariationPercentage float64
	VariationAmount     float64
	AssetCurrentPrice   float64
}

func NewConsolidatorItem(grouper string, quantity int, avgPrice float64, total float64, percentage float64, details string) ConsolidatorItem {
	return ConsolidatorItem{
		grouper,
		quantity,
		avgPrice,
		total,
		total,
		total,
		percentage,
		details,
		0,
		0,
		0,
	}
}

func (ci *ConsolidatorItem) Reset() *ConsolidatorItem {
	ci.TotalQuantity = 0
	ci.PaidCost = 0.0
	ci.AverageAmount = 0.0
	ci.AveragePrice = 0.0
	ci.WalletPercentage = 0.0
	return ci
}

func (ci ConsolidatorItem) Add(transaction Transaction) ConsolidatorItem {
	ci.PaidCost += transaction.TotalWithoutTaxes()
	ci.TotalQuantity += transaction.Quantity
	if ci.TotalQuantity <= 0 {
		ci.Reset()
		return ci
	}

	ci.AveragePrice = roundFloat(ci.PaidCost / float64(ci.TotalQuantity))
	ci.AverageAmount = roundFloat(float64(ci.TotalQuantity) * ci.AveragePrice)
	return ci
}

func (ci *ConsolidatorItem) PercentageOf(amount float64) *ConsolidatorItem {
	percentage := ci.PaidCost * 100 / amount
	if percentage < 0 {
		percentage = 0.0
	}
	ci.WalletPercentage = percentage
	return ci
}

func (ci *ConsolidatorItem) CurrentPrice(price float64) *ConsolidatorItem {
	totalForPrice := roundFloat(price * float64(ci.TotalQuantity))
	ci.AverageAmount = roundFloat(ci.AveragePrice * float64(ci.TotalQuantity))
	variationAmount := roundFloat(totalForPrice - ci.AverageAmount)
	variationPercentage := 0.0
	if variationAmount != 0 {
		variationPercentage = roundFloat((variationAmount * 100) / ci.AverageAmount)
	}

	ci.CurrentAmount = totalForPrice
	ci.AssetCurrentPrice = price
	ci.VariationPercentage = variationPercentage
	ci.VariationAmount = variationAmount
	return ci
}

func (ci ConsolidatorItem) HasGrouper(kind string) bool {
	return ci.Grouper == kind
}

func (ci ConsolidatorItem) HasDetails(detail string) bool {
	return ci.Details == detail
}

func ConsolidateByKind(wallet Wallet) map[string]ConsolidatorItem {
	walletTotal := wallet.Total()
	byKind := make(map[string]ConsolidatorItem)

	for _, consolidated := range wallet.Consolidation {
		assetKind := consolidated.Details
		byKindConsolidator, existsInMap := byKind[assetKind]

		if !existsInMap {
			byKindConsolidator = NewConsolidatorItem(assetKind, 0, 0.0, 0.0, 0.0, "")
		}

		byKindConsolidator.TotalQuantity += consolidated.TotalQuantity
		byKindConsolidator.AveragePrice = byKindConsolidator.AverageAmount / float64(byKindConsolidator.TotalQuantity)
		byKindConsolidator.AverageAmount += consolidated.AverageAmount
		byKindConsolidator.CurrentAmount += consolidated.CurrentAmount
		byKindConsolidator.PaidCost += consolidated.PaidCost
		byKindConsolidator.VariationAmount += consolidated.VariationAmount
		variationPercent := 0.0
		if byKindConsolidator.CurrentAmount != 0 {
			variationPercent = roundFloat((byKindConsolidator.VariationAmount * 100) / byKindConsolidator.CurrentAmount)
		}

		byKindConsolidator.VariationPercentage = variationPercent
		byKindConsolidator.PercentageOf(walletTotal)
		byKind[assetKind] = byKindConsolidator
	}

	return byKind
}

func ConsolidateByAsset(wallet Wallet) map[string]ConsolidatorItem {
	consolidationMap := make(map[string]ConsolidatorItem)
	tickers := []string{}

	for _, transaction := range wallet.Transactions {
		asset := transaction.Asset
		ticker := asset.Ticker
		consolidator, alreadyOnMap := consolidationMap[ticker]

		if !alreadyOnMap {
			tickers = append(tickers, ticker)
			consolidator = NewConsolidatorItem(ticker, 0, 0.0, 0.0, 0.0, asset.Kind)
		}

		consolidator = consolidator.Add(transaction)
		consolidationMap[ticker] = consolidator
	}

	tickerService := brapi.NewTickerService()
	currentPrices, err := tickerService.GetPricesByCodes(tickers)

	if err != nil {
		fmt.Println(err)
	}

	walletTotal := 0.0
	for assetTicker, consolidation := range consolidationMap {
		currentPrice := currentPrices[assetTicker]
		consolidation.CurrentPrice(currentPrice)
		walletTotal += consolidation.CurrentAmount
		consolidationMap[assetTicker] = consolidation
	}

	for assetTicker, consolidation := range consolidationMap {
		consolidation.PercentageOf(walletTotal)
		consolidationMap[assetTicker] = consolidation
	}

	return consolidationMap
}

func roundFloat(amount float64) float64 {
	ratio := math.Pow(10, float64(2))
	return math.Round(amount*ratio) / ratio
}
