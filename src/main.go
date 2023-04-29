package main

import (
	"investment-warlock/importers"
)

func main() {
	// wallet := investor.BuildWalletFromJsonFile("wallet.json")
	wallet, _ := importers.ImportFromB3Csv("transactions.csv")
	DisplayMenu(wallet)
}
