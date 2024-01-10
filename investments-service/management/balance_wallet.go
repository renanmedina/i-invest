package management

import (
	"fmt"
	"github.com/renanmedina/investment-warlock/investments-service/utils"
	"github.com/jedib0t/go-pretty/v6/table"
)

func BalanceWallet(wallet Wallet) {
	suggestions := setupBalancer(wallet)
	fmt.Println("")
	fmt.Println("Resultado do balanceamento:")
	fmt.Println("")
	fmt.Print("Total investido na carteira: R$ ", wallet.Total())
	fmt.Println("")
	printBalancingSummary(suggestions)
}

func setupBalancer(wallet Wallet) []BalanceSuggestion {
	utils.ClearConsole()
	fmt.Println("===========================================================")
	fmt.Println("            Balancear Carteira de investimento             ")
	fmt.Println("===========================================================")
	fmt.Print("% em ações: ")
	stock_percents := utils.ReadFloat()
	fmt.Print("% em Fundos Imobiliários: ")
	fii_percents := utils.ReadFloat()
	fmt.Print("% em Renda fixa: ")
	fixed_income_percents := utils.ReadFloat()
	setup := MakeBalanceSetup(fii_percents, stock_percents, fixed_income_percents)
	fmt.Println("Balanceando, aguarde ....")
	suggestions := BalanceWalletByAssetType(wallet, setup)
	return suggestions
}

func printBalancingSummary(suggestions []BalanceSuggestion) {
	AssetKindLegend := map[string]string{
		"stock":        "Ações",
		"real_state":   "Fundos Imobiliários",
		"fixed_income": "Renda Fixa",
	}

	writer := utils.NewTableWriter()
	writer.AppendHeader(table.Row{"Tipo de ativo", "% na carteira", "R$ na carteira", "% alvo", "Operação", "R$ Valor"})
	for _, suggestion := range suggestions {
		writer.AppendRow([]interface{}{
			AssetKindLegend[suggestion.AssetKind],
			utils.PercentageFormat(suggestion.CurrentPercentage),
			utils.CurrencyFormat(suggestion.CurrentTotal),
			utils.PercentageFormat(suggestion.TargetPercentage),
			suggestion.Operation,
			utils.CurrencyFormat(suggestion.OperationAmount),
		})
	}
	writer.Render()
}
