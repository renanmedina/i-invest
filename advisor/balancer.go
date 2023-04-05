package advisor

import (
	"fmt"
	"investment-warlock/investor"
)

type BalanceSetup struct {
	fii        float64
	acoes      float64
	renda_fixa float64
}

func MakeSetup(fii float64, acoes float64, renda_fixa float64) BalanceSetup {
	return BalanceSetup{fii, acoes, renda_fixa}
}

func BalanceWallet(wallet investor.Wallet, setup BalanceSetup) {
	fmt.Println("Running wallet balancer")
}
