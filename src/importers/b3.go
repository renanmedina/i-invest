package importers

import (
	"encoding/csv"
	"investment-warlock/investor"
	"os"
	"strconv"
	"strings"
)

type CsvTransationItem struct {
	Date      string
	AssetType string
	Ticker    string
	Quantity  int
	Price     float64
}

func ImportFromB3Csv(filepath string) (investor.Wallet, error) {
	csvFile, err := os.Open(filepath)
	transactions := []investor.Transaction{}

	if err != nil {
		return investor.Wallet{}, err
	}

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	transactionsData, err := csvReader.ReadAll()

	if err != nil {
		return investor.Wallet{}, err
	}

	for rowIndex, line := range transactionsData {
		// ignore header
		if rowIndex > 0 {
			record := parseTransactionLine(line)

			transaction := investor.NewTransaction(
				record.AssetType,
				record.Ticker,
				record.Price,
				record.Quantity,
				0.0,
				record.Date,
			)

			transactions = append(transactions, transaction)
		}
	}

	wallet := investor.NewWallet("1", "Wallet de testes", "Renan Medina", transactions).Consolidate()
	return wallet, nil
}

func parseTransactionLine(csvLine []string) CsvTransationItem {
	assetType := "stock"
	tickerId := csvLine[5][len(csvLine[5])-2:]
	if csvLine[2] == "Mercado à Vista" && tickerId == "11" {
		assetType = "real_state"
	}

	quantity, _ := strconv.Atoi(csvLine[6])
	replacedPrice := strings.ReplaceAll(strings.ReplaceAll(csvLine[7], "R$", ""), " ", "")
	price, _ := strconv.ParseFloat(replacedPrice, 64)

	if csvLine[1] == "Venda" {
		quantity *= -1
	}

	return CsvTransationItem{
		csvLine[0],
		assetType,
		csvLine[5],
		quantity,
		price,
	}
}
