package wallets

import (
	"fmt"
	"github.com/renanmedina/investment-warlock/backend/market/brapi"
	"github.com/renanmedina/investment-warlock/backend/utils"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

func SearchAsset(wallet Wallet) {
	fmt.Println("===========================================================")
	fmt.Println("              Consulta de Ativos da bolsa B3               ")
	fmt.Println("===========================================================")
	fmt.Print("Informe o codigo do ativo (passe separado por virgula para multipos): ")
	tickerCodes := strings.Split(utils.ReadLine(), ",")
	fmt.Println("Buscando, aguarde ....")
	service := brapi.NewTickerService()
	tickers, err := service.GetByCodes(tickerCodes)

	if err != nil {
		fmt.Println(err)
		return
	}

	printMarketTickers(tickers, wallet)
}

func printMarketTickers(tickers []brapi.Ticker, wallet Wallet) {
	writer := utils.NewTableWriter()
	writer.AppendHeader(table.Row{"Ativo", "Nome", "Preço", "Ultimo Fechamento", "R$ na carteira", "% na carteira"})
	for _, ticker := range tickers {
		consolidated, _ := wallet.GetConsolidation(ticker.Code)
		writer.AppendRow([]interface{}{
			ticker.Code,
			ticker.Name,
			utils.CurrencyFormat(ticker.Price),
			utils.CurrencyFormat(ticker.LastClosePrice),
			utils.CurrencyFormat(consolidated.AverageAmount),
			utils.PercentageFormat(consolidated.WalletPercentage),
		})
	}
	writer.Render()
}
