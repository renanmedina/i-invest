package use_cases

import (
	"fmt"
	"investment-warlock/investor"
	"investment-warlock/utils"

	"github.com/jedib0t/go-pretty/v6/table"
)

func DisplayConsolidation(wallet investor.Wallet) {
	printConsolidationByKind(wallet)
	displayConsolidationSubmenu(wallet)
}

func printConsolidationByKind(wallet investor.Wallet) {
	utils.PrintWalletHeader(wallet)
	writer := utils.NewTableWriter()
	writer.AppendHeader(table.Row{"Tipo", "Quantidade", "Preço Médio", "Total", "% Carteira"})
	consolidation := investor.ConsolidateByKind(wallet)

	for _, consolidatedGroup := range consolidation {
		assetType := utils.TranslateKind(consolidatedGroup.Grouper)
		averagePrice := utils.CurrencyFormat(consolidatedGroup.AveragePrice)
		total := utils.CurrencyFormat(consolidatedGroup.TotalCost)
		percentage := utils.PercentageFormat(consolidatedGroup.WalletPercentage)

		writer.AppendRow([]interface{}{
			assetType,
			consolidatedGroup.TotalQuantity,
			averagePrice,
			total,
			percentage,
		})
	}

	writer.SortBy([]table.SortBy{
		{Name: "% Carteira", Mode: table.DscNumeric},
		{Name: "Total", Mode: table.DscNumeric},
		{Name: "Quantidade", Mode: table.DscNumeric},
	})

	writer.Render()
}

func displayConsolidationSubmenu(wallet investor.Wallet) {
	fmt.Println("")
	fmt.Println("1 - Exibir consolidação por ativo")
	fmt.Println("2 - Exibir consolidação setor de ativo")
	fmt.Println("3 - Voltar")
	fmt.Print("Opção: ")
	option := utils.ReadOption()
	switch option {
	case 1:
		printConsolidationByAsset(wallet)
		break
	case 2:
		fmt.Println("Ainda não implementado")
	}
}

func printConsolidationByAsset(wallet investor.Wallet) {
	utils.PrintWalletHeader(wallet)
	writer := utils.NewTableWriter()
	writer.AppendHeader(table.Row{"Ativo", "Tipo", "Quantidade", "Preço Médio", "Total", "% Carteira"})
	for _, consolidatedAsset := range wallet.Consolidation {
		assetType := utils.TranslateKind(consolidatedAsset.Details)
		quantity := consolidatedAsset.TotalQuantity
		averagePrice := utils.CurrencyFormat(consolidatedAsset.AveragePrice)
		total := utils.CurrencyFormat(consolidatedAsset.TotalCost)
		percentage := utils.PercentageFormat(consolidatedAsset.WalletPercentage)

		writer.AppendRow([]interface{}{consolidatedAsset.Grouper, assetType, quantity, averagePrice, total, percentage})
	}

	writer.SortBy([]table.SortBy{
		{Name: "% Carteira", Mode: table.DscNumeric},
		{Name: "Total", Mode: table.DscNumeric},
		{Name: "Quantidade", Mode: table.DscNumeric},
	})

	writer.Render()
}
