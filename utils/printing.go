package utils

import (
	"fmt"
	"investment-warlock/investor"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintConsolidation(wallet investor.Wallet) {
	fmt.Println("===========================================================")
	fmt.Println("Wallet: ", wallet.Name)
	fmt.Println("Cliente: ", wallet.Client.Name)
	fmt.Println("Total investido: R$ ", wallet.Total())
	writer := table.NewWriter()
	writer.SetOutputMirror(os.Stdout)
	writer.AppendHeader(table.Row{"Ativo", "Quantidade", "Preço Médio", "Total"})
	for _, consolidatedAsset := range wallet.Consolidation {
		asset := consolidatedAsset.Asset
		quantity := consolidatedAsset.TotalQuantity
		averagePrice := consolidatedAsset.AveragePrice
		total := consolidatedAsset.TotalCost

		writer.AppendRow([]interface{}{asset.Ticker, quantity, averagePrice, total})
	}

	writer.Render()
}
