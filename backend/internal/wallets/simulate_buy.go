package wallets

import (
	"fmt"
	"github.com/renanmedina/investment-warlock/backend/utils"
)

func SimulateBuy(wallet Wallet) {
	utils.ClearConsole()
	fmt.Println("===========================================================")
	fmt.Println("                Simular aporte de compra                   ")
	fmt.Println("===========================================================")
	fmt.Print("Informe o valor do aporte R$: ")
}
