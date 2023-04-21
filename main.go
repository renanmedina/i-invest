package main

import (
	"investment-warlock/investor"
	"investment-warlock/utils"
)

func main() {
	// wallet := investor.BuildWalletFromJsonFile("wallet.json")
	wallet, _ := investor.ImportFromCsv("transactions.csv")
	utils.DisplayMenu(wallet)
}
