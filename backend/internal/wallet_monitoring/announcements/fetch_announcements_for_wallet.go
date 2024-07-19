package announcements

import (
	"time"

	"github.com/renanmedina/i-invest/internal/wallets"
)

type FetchAnnouncementsForWallet struct {
	allWallets wallets.WalletRepository
}

func NewFetchAnnouncementsForWallet() *FetchAnnouncementsForWallet {
	return &FetchAnnouncementsForWallet{
		*wallets.NewWalletRepository(),
	}
}

func (uc *FetchAnnouncementsForWallet) execute(walletId string) {
	wallet, err := uc.allWallets.GetById(walletId)

	if err != nil {
		panic(err)
	}

	fetchUseCase := NewFetchCompanyNewAnnouncements()
	today := time.Now()

	for _, tickerCode := range wallet.GetTickerCodes() {
		fetchUseCase.Execute(tickerCode, today.Year())
	}
}
