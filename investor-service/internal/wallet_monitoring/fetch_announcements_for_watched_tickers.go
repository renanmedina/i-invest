package wallet_monitoring

import (
	"time"

	"github.com/renanmedina/i-invest/internal/wallet_monitoring/announcements"
	"github.com/renanmedina/i-invest/internal/wallet_monitoring/watchlists"
)

type FetchAnnouncementsForWatchedTickers struct {
	allWatchedTickers watchlists.WatchedAssetRepository
}

func NewFetchAnnouncementsForWatchedTickers() FetchAnnouncementsForWatchedTickers {
	return FetchAnnouncementsForWatchedTickers{
		*watchlists.NewWatchedAssetRepository(),
	}
}

func (uc *FetchAnnouncementsForWatchedTickers) Execute() {
	tickers := uc.allWatchedTickers.GetUniqTickersWithAnnouncementsEnabled()
	fetchUseCase := announcements.NewFetchCompanyNewAnnouncements()

	for _, tickerCode := range tickers {
		go fetchUseCase.Execute(tickerCode, time.Now().Year())
	}
}
