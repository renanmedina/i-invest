package utils

import (
	"fmt"
	"investment-warlock/investor"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintWalletHeader(wallet investor.Wallet) {
	fmt.Println("===========================================================")
	fmt.Println("Wallet: ", wallet.Name)
	fmt.Println("Cliente: ", wallet.Client.Name)
	fmt.Println("Total da carteira: ", CurrencyFormat(wallet.Total()))
}

func NewTableWriter() table.Writer {
	writer := table.NewWriter()
	writer.SetOutputMirror(os.Stdout)
	return writer
}
