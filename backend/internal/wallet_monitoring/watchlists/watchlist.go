package watchlists

import (
	"github.com/google/uuid"
	"github.com/renanmedina/investment-warlock/internal/wallets/b3"
)

type Watchlist struct {
	Id        uuid.UUID `json:"id"`
	userId    uuid.UUID
	Name      string `json:"name"`
	assets    []WatchedAsset
	Persisted bool
}

func NewEmptyWatchlist(userId uuid.UUID) Watchlist {
	return Watchlist{
		Id:     uuid.New(),
		Name:   "New Watchlist",
		userId: userId,
	}
}

func (w *Watchlist) assetsByCodeMap() map[string]string {
	byAssetCode := make(map[string]string)

	for _, watchedAsset := range w.assets {
		byAssetCode[watchedAsset.TickerCode] = watchedAsset.TickerCode
	}

	return byAssetCode
}

func (w *Watchlist) AddAsset(asset WatchedAsset) *Watchlist {
	w.assets = append(w.assets, asset)
	return w
}

func (w *Watchlist) importFromB3(items *[]b3.B3SummaryReportItem) *Watchlist {
	currentWatchedItems := w.assetsByCodeMap()

	for _, item := range *items {
		_, isWatched := currentWatchedItems[item.TickerCode]

		if isWatched {
			continue
		}

		w.AddAsset(NewWatchedAsset(item.TickerCode, item.Type))
	}

	return w
}
