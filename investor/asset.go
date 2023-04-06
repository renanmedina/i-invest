package investor

type Asset struct {
	Ticker string
	Price  float64
	Kind   string
}

type ConsolidatedAsset struct {
	Asset         Asset
	TotalCost     float64
	TotalQuantity int
	AveragePrice  float64
}

func (ca ConsolidatedAsset) Add(transaction Transaction) ConsolidatedAsset {
	ca.TotalCost += transaction.TotalWithoutTaxes()
	ca.TotalQuantity += transaction.Quantity
	ca.AveragePrice = ca.TotalCost / float64(ca.TotalQuantity)
	return ca
}
