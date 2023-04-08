package main

import (
	"investment-warlock/investor"
	"investment-warlock/utils"
)

func main() {
	wallet := investor.BuildWalletFromJsonFile("wallet.json")
	utils.DisplayMenu(wallet)
}
