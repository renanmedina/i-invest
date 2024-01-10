package use_cases

import (
	"fmt"
	"github.com/renanmedina/investment-warlock/investments-service/investor"
	"github.com/renanmedina/investment-warlock/investments-service/utils"
)

func SimulateBuy(wallet investor.Wallet) {
	utils.ClearConsole()
	fmt.Println("===========================================================")
	fmt.Println("                Simular aporte de compra                   ")
	fmt.Println("===========================================================")
	fmt.Print("Informe o valor do aporte R$: ")
}
