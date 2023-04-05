package main

import (
	"fmt"
	"investment-warlock/advisor"
	"investment-warlock/investor"
)

func main() {
	wallet := investor.BuildWalletFromJsonFile("wallet.json")
	advisor.BalanceWallet(wallet, advisor.MakeSetup(40, 30, 30))
	fmt.Println(wallet.Total())
	fmt.Println(wallet)
}
