package utils

import (
	"fmt"
	"github.com/renanmedina/investment-warlock/investments-service/investor"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintWalletHeader(wallet investor.Wallet) {
	fmt.Println("===========================================================")
	fmt.Println("Wallet: ", wallet.Name)
	fmt.Println("Cliente: ", wallet.Client.Name)
	fmt.Println("Patrimonio atual: ", CurrencyFormat(wallet.Total()))
	fmt.Println("Investimento realizado: ", CurrencyFormat(wallet.TotalInvested()))
	fmt.Println("% variação: ", PercentageFormat(wallet.VariationPercentage()))
}

func NewTableWriter() table.Writer {
	writer := table.NewWriter()
	writer.SetOutputMirror(os.Stdout)
	return writer
}
