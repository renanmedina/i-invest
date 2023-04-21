package utils

import (
	"bufio"
	"fmt"
	"investment-warlock/advisor"
	"investment-warlock/investor"
	"investment-warlock/market/brapi"
	"os"
	"os/exec"
	"strconv"
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

func readOption() uint64 {
	option, _ := strconv.ParseUint(readLine(), 10, 64)
	return option
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
	clear()
	switch option {
	case "1":
		PrintTransactions(wallet)
	case "2":
		PrintConsolidationByKind(wallet)
		displayConsolidationDetails(wallet)
		// PrintConsolidation(wallet)
	case "3":
		setupBalancer(wallet)
	case "5":
		displayMarketAsset(wallet)
	}

	fmt.Println("Pressione qualquer tecla para continuar ....")
	readLine()
}

func displayMarketAsset(wallet investor.Wallet) {
	fmt.Println("===========================================================")
	fmt.Println("              Consulta de Ativos da bolsa B3               ")
	fmt.Println("===========================================================")
	fmt.Print("Informe o codigo do ativo (passe separado por virgula para multipos): ")
	tickerCodes := strings.Split(readLine(), ",")
	fmt.Println("Buscando, aguarde ....")
	service := brapi.NewTickerService()
	tickers, err := service.GetByCodes(tickerCodes)

	if err != nil {
		fmt.Println(err)
		return
	}

	PrintMarketTicker(tickers, wallet)
}

func setupBalancer(wallet investor.Wallet) {
	clear()
	fmt.Println("===========================================================")
	fmt.Println("            Balancear Carteira de investimento             ")
	fmt.Println("===========================================================")
	fmt.Print("% em ações: ")
	stock_percents, _ := strconv.ParseFloat(readLine(), 32)
	fmt.Print("% em Fundos Imobiliários: ")
	fii_percents, _ := strconv.ParseFloat(readLine(), 32)
	fmt.Print("% em Renda fixa: ")
	fixed_income_percents, _ := strconv.ParseFloat(readLine(), 32)
	setup := advisor.MakeBalanceSetup(fii_percents, stock_percents, fixed_income_percents)
	// setup := advisor.MakeBalanceSetup(40, 30, 30)
	suggestions := advisor.BalanceWalletByAssetType(wallet, setup)
	fmt.Println("Balanceando, aguarde ....")
	fmt.Println("")
	fmt.Println("Resultado do balanceamento:")
	fmt.Println("")
	fmt.Print("Total investido na carteira: R$ ", wallet.Total())
	fmt.Println("")
	PrintBalancingSummary(suggestions)
}

func displayConsolidationDetails(wallet investor.Wallet) int {
	fmt.Println("")
	fmt.Println("1 - Exibir consolidação por ativo")
	fmt.Println("2 - Exibir consolidação setor de ativo")
	fmt.Println("3 - Voltar")
	fmt.Print("Opção: ")
	option := readOption()
	switch option {
	case 1:
		PrintConsolidation(wallet)
		return 0
	case 2:
		fmt.Println("Ainda não implementado")
	}

	return 0
}
