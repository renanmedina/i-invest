package main

import (
	"investment-warlock/advisor"
	"investment-warlock/investor"
	"investment-warlock/utils"
)

func main() {
	wallet := investor.BuildWalletFromJsonFile("wallet.json")
	advisor.BalanceWallet(wallet, advisor.MakeSetup(40, 30, 30))
	utils.PrintConsolidation(wallet)
}
