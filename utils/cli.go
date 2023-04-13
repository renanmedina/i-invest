package utils

import (
	"bufio"
	"fmt"
	"investment-warlock/investor"
	"investment-warlock/market"
	"os"
	"os/exec"
	"strings"
)

func clear() {
	command := exec.Command("clear")
	command.Stdout = os.Stdout
	command.Run()
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	selected_option, _ := reader.ReadString('\n')
	selected_option = strings.TrimSpace(selected_option)
	return selected_option
}

func DisplayMenu(wallet investor.Wallet) {
	exit_action := "8"
	for {
		clear()
		fmt.Println("===========================================================")
		fmt.Println("Olá, sr. investidor o que deseja fazer?")
		fmt.Println("1 - Listar transações")
		fmt.Println("2 - Exibir carteira consolidada")
		fmt.Println("3 - Balancear carteira")
		fmt.Println("4 - Simular aportes")
		fmt.Println("5 - Consultar ativo")
		fmt.Println("6 - Ranking de ativos")
		fmt.Println("7 - Dividendos")
		fmt.Println(exit_action, "- Sair")
		fmt.Println("===========================================================")

		fmt.Print("Informe a sua opção: ")
		selected_option := readLine()
		if selected_option == exit_action {
			break
		}

		executeAction(selected_option, wallet)
	}
}

func executeAction(option string, wallet investor.Wallet) {
	switch option {
	case "1":
		PrintTransactions(wallet)
	case "2":
		PrintConsolidation(wallet)
	case "5":
		displayMarketAsset()
	}

	fmt.Println("Pressione qualquer tecla para continuar ....")
	readLine()
}

func displayMarketAsset() {
	clear()
	fmt.Print("Informe o codigo do ativo: ")
	tickerCode := readLine()
	fmt.Println("Buscando ....")
	service := market.NewTicketService()
	ticker, err := service.GetByCode(tickerCode)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("===========================================================")
	fmt.Printf("Ativo: %s \r\n", ticker.Code)
	fmt.Printf("Nome: %s \r\n", ticker.Name)
	fmt.Printf("Preço: %s \r\n", currencyFormat(ticker.Price))
	fmt.Printf("Preço de ultimo fechamento: %s \r\n", currencyFormat(ticker.LastClosePrice))
	fmt.Println("===========================================================")
}
