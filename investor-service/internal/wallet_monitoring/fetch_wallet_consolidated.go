package wallet_monitoring

import (
	"fmt"

	B3 "github.com/renanmedina/i-invest/internal/market/b3"
	"github.com/renanmedina/i-invest/internal/wallets"
	"github.com/renanmedina/i-invest/utils"
)

type FetchWalletConsolidated struct {
	logger         *utils.ApplicationLogger
	allWallets     *wallets.WalletRepository
	walletsService *B3.WalletService
}

func NewFetchWalletConsolidated(sourceService *B3.WalletService) *FetchWalletConsolidated {
	return &FetchWalletConsolidated{
		logger:         utils.GetApplicationLogger(),
		allWallets:     wallets.NewWalletRepository(),
		walletsService: sourceService,
	}
}

func (uc *FetchWalletConsolidated) Execute(walletId string, snapshotDate string) error {
	_, err := uc.allWallets.GetById(walletId)

	if err != nil {
		return err
	}

	uc.logger.Info(fmt.Sprintf("Fetching wallet consolidated from source for snapshot date %s", snapshotDate))
	// _, err := uc.walletsService.GetConsolidatedSnapshotByDate(snapshotDate)

	// if err != nil {
	// 	return err
	// }

	// uc.logger.Info(fmt.Sprintf("Replacing consolidated from source for snapshot date %s", snapshotDate))
	// uc.logger.Info(fmt.Sprintf("Saving walet %s", wallet.Id))
	// uc.allWallets.Save(*wallet)

	return nil
}
