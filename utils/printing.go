package utils

import (
	"fmt"
	"investment-warlock/investor"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func printWalletHeader(wallet investor.Wallet) {
	fmt.Println("===========================================================")
	fmt.Println("Wallet: ", wallet.Name)
	fmt.Println("Cliente: ", wallet.Client.Name)
	fmt.Println("Total investido: R$ ", wallet.Total())
}

func makeWriter() table.Writer {
	writer := table.NewWriter()
	writer.SetOutputMirror(os.Stdout)
	return writer
}

func PrintConsolidation(wallet investor.Wallet) {
	printWalletHeader(wallet)
	writer := makeWriter()
	writer.AppendHeader(table.Row{"Ativo", "Tipo", "Quantidade", "Preço Médio", "Total", "% Carteira"})
	for _, consolidatedAsset := range wallet.Consolidation {
		asset := consolidatedAsset.Asset
		assetType := translateKind(asset.Kind)
		quantity := consolidatedAsset.TotalQuantity
		averagePrice := currencyFormat(consolidatedAsset.AveragePrice)
		total := currencyFormat(consolidatedAsset.TotalCost)
		percentage := percentageFormat(consolidatedAsset.WalletPercentage)

		writer.AppendRow([]interface{}{asset.Ticker, assetType, quantity, averagePrice, total, percentage})
	}

	writer.SortBy([]table.SortBy{
		{Name: "% Carteira", Mode: table.DscNumeric},
		{Name: "Total", Mode: table.DscNumeric},
		{Name: "Quantidade", Mode: table.DscNumeric},
	})

	writer.Render()
}

func PrintTransactions(wallet investor.Wallet) {
	printWalletHeader(wallet)
	writer := makeWriter()
	writer.AppendHeader(table.Row{"Data", "Ativo", "Preço", "Quantidade", "Total"})
	for _, transaction := range wallet.Transactions {
		date := transaction.TransactionDate
		asset := transaction.Asset
		quantity := transaction.Quantity
		price := currencyFormat(asset.Price)
		total := currencyFormat(transaction.Total())
		writer.AppendRow([]interface{}{date, asset.Ticker, price, quantity, total})
	}

	writer.SortBy([]table.SortBy{{Name: "Data", Mode: table.Dsc}})
	writer.Render()
}
