package use_cases

import (
	"fmt"
	"investment-warlock/investor"
	"investment-warlock/utils"
)

func SimulateBuy(wallet investor.Wallet) {
	utils.ClearConsole()
	fmt.Println("===========================================================")
	fmt.Println("                Simular aporte de compra                   ")
	fmt.Println("===========================================================")
	fmt.Print("Informe o valor do aporte R$: ")
}
