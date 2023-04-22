package main

import (
	"investment-warlock/investor"
)

func main() {
	// wallet := investor.BuildWalletFromJsonFile("wallet.json")
	wallet, _ := investor.ImportFromCsv("transactions.csv")
	DisplayMenu(wallet)
}
