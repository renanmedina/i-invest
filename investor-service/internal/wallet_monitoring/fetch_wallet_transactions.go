package wallet_monitoring

import (
	"fmt"

	B3 "github.com/renanmedina/i-invest/internal/market/b3"
	"github.com/renanmedina/i-invest/internal/wallets"
	"github.com/renanmedina/i-invest/utils"
)

type FetchWalletTransactions struct {
	logger         *utils.ApplicationLogger
	allWallets     *wallets.WalletRepository
	walletsService *B3.WalletService
}

func NewFetchWalletTransactions(sourceService *B3.WalletService) *FetchWalletTransactions {
	return &FetchWalletTransactions{
		logger:         utils.GetApplicationLogger(),
		allWallets:     wallets.NewWalletRepository(),
		walletsService: sourceService,
	}
}

func (uc *FetchWalletTransactions) Execute(walletId string, dateStart string, dateEnd string) error {
	wallet, err := uc.allWallets.GetById(walletId)

	if err != nil {
		return err
	}

	uc.logger.Info(fmt.Sprintf("Fetching transactions from source for period %s - %s", dateStart, dateEnd))
	negotiations, err := uc.walletsService.GetNegotiationsByPeriod(dateStart, dateEnd)

	if err != nil {
		return err
	}

	translatedTransactions := uc.translateFromSource(negotiations)
	uc.logger.Info(fmt.Sprintf("Adding transactions from source for period %s - %s to wallet", dateStart, dateEnd))
	wallet.AddTransactions(translatedTransactions)
	uc.logger.Info(fmt.Sprintf("Saving wallet %s", wallet.Id))
	uc.allWallets.Save(*wallet)

	return nil
}

func (uc *FetchWalletTransactions) translateFromSource(negotiations []B3.NegotiationDayItem) []wallets.Transaction {
	list := make([]wallets.Transaction, 0)
	for _, dayItem := range negotiations {
		for _, transactionItem := range dayItem.Negotiations {
			var kind string
			translated := wallets.Transaction{
				Id:       "",
				Kind:     kind,
				Quantity: int(transactionItem.Quantity),
				Asset: wallets.Asset{
					Kind:   kind,
					Price:  transactionItem.UnitPrice,
					Ticker: transactionItem.TickerCode,
				},
				TransactionDate: transactionItem.Date,
			}
			list = append(list, translated)
		}
	}

	return list
}
