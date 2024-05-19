package watchlists

import (
	"github.com/Masterminds/squirrel"
	"github.com/renanmedina/investment-warlock/utils"
)

type WatchedAssetRepository struct {
	db *utils.DatabaseAdapdater
}

func NewWatchedAssetRepository() *WatchedAssetRepository {
	return &WatchedAssetRepository{
		utils.GetDatabase(),
	}
}

func (r *WatchedAssetRepository) GetUniqTickersWithAnnouncementsEnabled() []string {
	query, _ := squirrel.
		Select("DISTINCT ticker_code").
		From(ASSETS_TABLE_NAME).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(r.db.GetConnection()).
		Query()

	tickerCodes := make([]string, 0)
	for query.Next() {
		var tickerCode string
		query.Scan(&tickerCode)
		tickerCodes = append(tickerCodes, tickerCode)
	}

	return tickerCodes
}
