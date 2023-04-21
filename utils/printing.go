package utils

import (
	"fmt"
	"investment-warlock/advisor"
	"investment-warlock/investor"
	"investment-warlock/market/brapi"
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
		assetType := translateKind(consolidatedAsset.Details)
		quantity := consolidatedAsset.TotalQuantity
		averagePrice := currencyFormat(consolidatedAsset.AveragePrice)
		total := currencyFormat(consolidatedAsset.TotalCost)
		percentage := percentageFormat(consolidatedAsset.WalletPercentage)

		writer.AppendRow([]interface{}{consolidatedAsset.Grouper, assetType, quantity, averagePrice, total, percentage})
	}

	writer.SortBy([]table.SortBy{
		{Name: "% Carteira", Mode: table.DscNumeric},
		{Name: "Total", Mode: table.DscNumeric},
		{Name: "Quantidade", Mode: table.DscNumeric},
	})

	writer.Render()
}

func PrintConsolidationByKind(wallet investor.Wallet) {
	printWalletHeader(wallet)
	writer := makeWriter()
	writer.AppendHeader(table.Row{"Tipo", "Quantidade", "Preço Médio", "Total", "% Carteira"})
	consolidation := investor.ConsolidateByKind(wallet)

	for _, consolidatedGroup := range consolidation {
		assetType := translateKind(consolidatedGroup.Grouper)
		averagePrice := currencyFormat(consolidatedGroup.AveragePrice)
		total := currencyFormat(consolidatedGroup.TotalCost)
		percentage := percentageFormat(consolidatedGroup.WalletPercentage)

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

func PrintMarketTicker(tickers []brapi.Ticker, wallet investor.Wallet) {
	writer := makeWriter()
	writer.AppendHeader(table.Row{"Ativo", "Nome", "Preço", "Ultimo Fechamento", "R$ na carteira", "% na carteira"})
	for _, ticker := range tickers {
		consolidated, _ := wallet.GetConsolidation(ticker.Code)
		writer.AppendRow([]interface{}{
			ticker.Code,
			ticker.Name,
			currencyFormat(ticker.Price),
			currencyFormat(ticker.LastClosePrice),
			currencyFormat(consolidated.TotalCost),
			percentageFormat(consolidated.WalletPercentage),
		})
	}
	writer.Render()
}

func PrintBalancingSummary(suggestions []advisor.BalanceSuggestion) {
	AssetKindLegend := map[string]string{
		"stock":        "Ações",
		"real_state":   "Fundos Imobiliários",
		"fixed_income": "Renda Fixa",
	}

	writer := makeWriter()
	writer.AppendHeader(table.Row{"Tipo de ativo", "% na carteira", "R$ na carteira", "% alvo", "Operação", "R$ Valor"})
	for _, suggestion := range suggestions {
		writer.AppendRow([]interface{}{
			AssetKindLegend[suggestion.AssetKind],
			percentageFormat(suggestion.CurrentPercentage),
			currencyFormat(suggestion.CurrentTotal),
			percentageFormat(suggestion.TargetPercentage),
			suggestion.Operation,
			currencyFormat(suggestion.OperationAmount),
		})
	}
	writer.Render()
}
