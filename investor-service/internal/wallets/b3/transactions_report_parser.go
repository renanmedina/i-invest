package b3

import (
	"strconv"
	"strings"

	"github.com/thedatashed/xlsxreader"
)

const (
	DEBIT_TRANSACTION_TYPE  = "Debito"
	CREDIT_TRANSACTION_TYPE = "Credito"
)

type B3TransationReportItem struct {
	Date      string
	Details   string
	AssetType string
	Ticker    string
	Quantity  float64
	Price     float64
	Total     float64
}

func ParseTransactionsReport(filepath string) ([]B3TransationReportItem, error) {
	xl, err := xlsxreader.OpenFile(filepath)
	transactions := make([]B3TransationReportItem, 0)

	if err != nil {
		return transactions, err
	}

	defer xl.Close()

	if err != nil {
		return transactions, err
	}

	headerSkipped := false
	for row := range xl.ReadRows(xl.Sheets[0]) {
		if !headerSkipped {
			headerSkipped = true
			continue
		}

		price, _ := strconv.ParseFloat(row.Cells[6].Value, 64)
		total, _ := strconv.ParseFloat(row.Cells[7].Value, 64)
		quantity, _ := strconv.ParseFloat(row.Cells[5].Value, 64)

		if row.Cells[0].Value == DEBIT_TRANSACTION_TYPE {
			quantity *= -1
			total *= -1
		}

		tickerId := row.Cells[3].Value
		tickerSplit := strings.Split(tickerId, "-")
		if len(tickerSplit) > 1 {
			tickerId = strings.TrimSpace(tickerSplit[0])
		}

		assetType := "stock"

		if tickerId[len(tickerId)-2:] == "11" {
			assetType = "real_state"
		}

		if len(tickerId) > 6 && tickerId[0:6] == "Tesouro" {
			assetType = "treasure"
		}

		transactionItem := B3TransationReportItem{
			Date:      strings.TrimSpace(row.Cells[1].Value),
			Details:   strings.TrimSpace(row.Cells[2].Value),
			AssetType: assetType,
			Ticker:    tickerId,
			Quantity:  quantity,
			Price:     price,
			Total:     total,
		}
		transactions = append(transactions, transactionItem)
	}

	return transactions, nil
}
