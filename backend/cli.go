package main

import (
	"fmt"
	"investment-warlock/investor"
	wallet_repository "investment-warlock/investor/repositories"
	"investment-warlock/use_cases"
	"investment-warlock/utils"
)

func DisplayMenu(wallet investor.Wallet) {
	exit_action := "9"
	for {
		// utils.ClearConsole()
		fmt.Println("===========================================================")
		fmt.Println("Olá, sr. investidor o que deseja fazer?")
		fmt.Println("1 - Listar transações")
		fmt.Println("2 - Exibir carteira consolidada")
		fmt.Println("3 - Balancear carteira")
		fmt.Println("4 - Simular aportes")
		fmt.Println("5 - Consultar ativo")
		fmt.Println("6 - Ranking de ativos")
		fmt.Println("7 - Dividendos")
		fmt.Println("8 - Salvar carteira")
		fmt.Println(exit_action, "- Sair")
		fmt.Println("===========================================================")

		fmt.Print("Informe a sua opção: ")
		selected_option := utils.ReadLine()
		if selected_option == exit_action {
			break
		}

		executeAction(selected_option, wallet)
	}
}

func executeAction(option string, wallet investor.Wallet) {
	utils.ClearConsole()
	switch option {
	case "1":
		use_cases.DisplayTransactions(wallet)
	case "2":
		use_cases.DisplayConsolidation(wallet)
	case "3":
		use_cases.BalanceWallet(wallet)
	case "4":
	case "5":
		use_cases.SearchAsset(wallet)
	case "8":
		wallet_repository.Save(wallet)
	}

	fmt.Println("Pressione qualquer tecla para continuar ....")
	utils.ReadLine()
}
