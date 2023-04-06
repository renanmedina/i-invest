package utils

import (
	"fmt"
	"investment-warlock/investor"
)

func PrintConsolidation(wallet investor.Wallet) {
	fmt.Println("========================================================================")
	fmt.Println("Wallet: ", wallet.Name)
	fmt.Println("Cliente: ", wallet.Client.Name)
	fmt.Println("Total investido: R$ ", wallet.Total())
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("| Ativo | \t Quantidade | \t Preço Médio | \t Total |")
	for _, consolidatedAsset := range wallet.Consolidation {
		asset := consolidatedAsset.Asset
		quantity := consolidatedAsset.TotalQuantity
		averagePrice := consolidatedAsset.AveragePrice
		total := consolidatedAsset.TotalCost
		fmt.Println(asset.Ticker, " | \t", quantity, " | \t", averagePrice, " | \t", total, " |")
	}
	fmt.Println("========================================================================")
}
