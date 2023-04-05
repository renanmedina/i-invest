package investor

type Transaction struct {
	Kind            string
	Quantity        int
	Taxes           float64
	Asset           Asset
	TransactionDate string
}

func (t Transaction) Total() float64 {
	return (t.Asset.Price * float64(t.Quantity)) + t.Taxes
}

func (t Transaction) TotalWithoutTaxes() float64 {
	return t.Asset.Price * float64(t.Quantity)
}
