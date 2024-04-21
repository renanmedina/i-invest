package b3

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"

	"github.com/renanmedina/investment-warlock/internal/wallets"
)

type CsvTransationItem struct {
	Date      string
	AssetType string
	Ticker    string
	Quantity  int
	Price     float64
}

func ParseB3CsvFile(filepath string) (wallets.Wallet, error) {
	csvFile, err := os.Open(filepath)
	transactions := []wallets.Transaction{}

	if err != nil {
		return wallets.Wallet{}, err
	}

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	transactionsData, err := csvReader.ReadAll()

	if err != nil {
		return wallets.Wallet{}, err
	}

	for rowIndex, line := range transactionsData {
		// ignore header
		if rowIndex > 0 {
			record := parseTransactionLine(line)

			transaction := wallets.NewTransaction(
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

	wallet := wallets.NewWallet("1", "Wallet de testes", "Renan Medina", transactions).Consolidate()
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
