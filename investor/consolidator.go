package investor

type ConsolidatorItem struct {
	Grouper          string
	TotalQuantity    int
	AveragePrice     float64
	TotalCost        float64
	WalletPercentage float64
	Details          string
}

func NewConsolidatorItem(grouper string, quantity int, avgPrice float64, total float64, percentage float64, details string) ConsolidatorItem {
	return ConsolidatorItem{
		grouper,
		quantity,
		avgPrice,
		total,
		percentage,
		details,
	}
}

func (ci ConsolidatorItem) Add(transaction Transaction) ConsolidatorItem {
	ci.TotalCost += transaction.TotalWithoutTaxes()
	ci.TotalQuantity += transaction.Quantity
	if ci.TotalQuantity <= 0 {
		ci.AveragePrice = 0.0
		return ci
	}

	ci.AveragePrice = ci.TotalCost / float64(ci.TotalQuantity)
	return ci
}

func (ci *ConsolidatorItem) PercentageOf(amount float64) *ConsolidatorItem {
	percentage := ci.TotalCost * 100 / amount
	ci.WalletPercentage = percentage
	return ci
}

func (ci ConsolidatorItem) HasKind(kind string) bool {
	return ci.Grouper == kind
}

func ConsolidateByKind(wallet Wallet) map[string]ConsolidatorItem {
	walletTotal := wallet.Total()
	byKind := make(map[string]ConsolidatorItem)

	for _, consolidated := range wallet.Consolidation {
		assetKind := consolidated.Details
		byKindConsolidator, existsInMap := byKind[assetKind]

		if !existsInMap {
			byKindConsolidator = ConsolidatorItem{assetKind, 0, 0.0, 0.0, 0.0, ""}
		}

		byKindConsolidator.TotalQuantity += consolidated.TotalQuantity
		byKindConsolidator.AveragePrice = byKindConsolidator.TotalCost / float64(byKindConsolidator.TotalQuantity)
		byKindConsolidator.TotalCost += consolidated.TotalCost
		byKindConsolidator.PercentageOf(walletTotal)
		byKind[assetKind] = byKindConsolidator
	}

	return byKind
}

func ConsolidateByAsset(wallet Wallet) map[string]ConsolidatorItem {
	consolidationMap := make(map[string]ConsolidatorItem)

	for _, transaction := range wallet.Transactions {
		asset := transaction.Asset
		ticker := asset.Ticker
		consolidator, alreadyOnMap := consolidationMap[ticker]

		if !alreadyOnMap {
			consolidator = NewConsolidatorItem(ticker, 0, 0.0, 0.0, 0.0, asset.Kind)
		}

		consolidator = consolidator.Add(transaction)
		consolidationMap[ticker] = consolidator
	}

	walletTotal := wallet.Total()
	for assetTicker, consolidation := range consolidationMap {
		consolidation.PercentageOf(walletTotal)
		consolidationMap[assetTicker] = consolidation
	}

	return consolidationMap
}
