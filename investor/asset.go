package investor

type Asset struct {
	Ticker string
	Price  float64
	Kind   string
}

type ConsolidatedAsset struct {
	Asset            Asset
	TotalCost        float64
	TotalQuantity    int
	AveragePrice     float64
	WalletPercentage float64
}

func (ca ConsolidatedAsset) Add(transaction Transaction) ConsolidatedAsset {
	ca.TotalCost += transaction.TotalWithoutTaxes()
	ca.TotalQuantity += transaction.Quantity
	ca.AveragePrice = ca.TotalCost / float64(ca.TotalQuantity)
	return ca
}

func (ca *ConsolidatedAsset) PercentageOf(amount float64) *ConsolidatedAsset {
	percentage := ca.TotalCost * 100 / amount
	ca.WalletPercentage = percentage
	return ca
}

func NewConsolidatedAsset(asset Asset, total float64, quantity int, price float64) ConsolidatedAsset {
	consolidated := ConsolidatedAsset{}
	consolidated.Asset = asset
	consolidated.TotalCost = total
	consolidated.TotalQuantity = quantity
	consolidated.AveragePrice = price
	return consolidated
}
