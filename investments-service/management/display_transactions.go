package management

import (
	"github.com/renanmedina/investment-warlock/investments-service/utils"

	"github.com/jedib0t/go-pretty/v6/table"
)

func DisplayTransactions(wallet Wallet) {
	wallet.PrintWalletHeader()
	writer := utils.NewTableWriter()
	writer.AppendHeader(table.Row{"Data", "Ativo", "Preço", "Quantidade", "Total"})
	for _, transaction := range wallet.Transactions {
		date := transaction.TransactionDate
		asset := transaction.Asset
		quantity := transaction.Quantity
		price := utils.CurrencyFormat(asset.Price)
		total := utils.CurrencyFormat(transaction.Total())
		writer.AppendRow([]interface{}{date, asset.Ticker, price, quantity, total})
	}

	writer.SortBy([]table.SortBy{{Name: "Data", Mode: table.Dsc}})
	writer.Render()
}
