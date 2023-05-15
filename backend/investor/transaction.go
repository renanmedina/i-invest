package investor

type Transaction struct {
	Id              string
	Kind            string
	Quantity        int
	Taxes           float64
	Asset           Asset
	TransactionDate string
}

func NewTransaction(assetKind string, ticker string, price float64, quantity int, taxes float64, date string) Transaction {
	asset := NewAsset(assetKind, ticker, price)
	transactionType := "compra"
	if quantity < 0 {
		transactionType = "venda"
	}
	return Transaction{Kind: transactionType, Quantity: quantity, Taxes: taxes, TransactionDate: date, Asset: asset}
}

func (t Transaction) Total() float64 {
	return (t.Asset.Price * float64(t.Quantity)) + t.Taxes
}

func (t Transaction) TotalWithoutTaxes() float64 {
	return t.Asset.Price * float64(t.Quantity)
}

func (t Transaction) Ticker() string {
	return t.Asset.Ticker
}

func (t Transaction) AssetPrice() float64 {
	return t.Asset.Price
}

func (t Transaction) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"transaction_date": t.TransactionDate,
		"transaction_type": t.Kind,
		"ticker":           t.Asset.Ticker,
		"asset_type":       t.Asset.Kind,
		"asset_price":      t.AssetPrice(),
		"quantity":         t.Quantity,
		"total":            t.Total(),
	}
}
