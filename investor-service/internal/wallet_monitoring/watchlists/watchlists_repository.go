package watchlists

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/renanmedina/i-invest/utils"
)

const TABLE_NAME = "watchlists"
const ASSETS_TABLE_NAME = "watchlists_assets"

type WatchlistsRepository struct {
	db *utils.DatabaseAdapdater
}

func NewWatchlistsRepository() *WatchlistsRepository {
	return &WatchlistsRepository{
		utils.GetDatabase(),
	}
}

func BuildWatchlistFromDb(watchlistDbRow squirrel.RowScanner, assetsDbRows *sql.Rows) (*Watchlist, error) {
	var watchlist Watchlist

	watchlistDbRow.Scan(
		&watchlist.Id,
		&watchlist.userId,
		&watchlist.Name,
	)

	if watchlist.Id.String() == "" {
		return nil, errors.New("can't find Watchlist")
	}

	for assetsDbRows.Next() {
		var assetCode, assetType string
		assetsDbRows.Scan(&assetCode, assetType)
		asset := NewWatchedAsset(assetCode, assetType)
		asset.Persisted = true
		watchlist.AddAsset(asset)
	}

	watchlist.Persisted = true
	return &watchlist, nil
}

func (r *WatchlistsRepository) GetById(watchlistId string) (*Watchlist, error) {
	query := squirrel.
		Select("*").
		From(TABLE_NAME).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": watchlistId}).
		RunWith(r.db.GetConnection())

	queryAssets, _ := squirrel.
		Select("ticker_code, ticker_type, configs").
		From(ASSETS_TABLE_NAME).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"watchlist_id": watchlistId}).
		RunWith(r.db.GetConnection()).Query()

	watchlist, err := BuildWatchlistFromDb(query.QueryRow(), queryAssets)

	if err != nil {
		return nil, err
	}

	return watchlist, nil
}

func (r *WatchlistsRepository) Save(watchlist *Watchlist) (*Watchlist, error) {
	var err error

	if watchlist.Persisted {
		_, err = r.db.Update(TABLE_NAME, map[string]interface{}{
			"label":      watchlist.Name,
			"updated_at": time.Now(),
		}, squirrel.Eq{"id": watchlist.Id})
	} else {
		_, err = r.db.Insert(TABLE_NAME, map[string]interface{}{
			"id":      watchlist.Id,
			"user_id": watchlist.userId,
			"label":   watchlist.Name,
		})
	}

	if err != nil {
		return nil, err
	}

	r.saveAssets(watchlist)
	return watchlist, nil
}

func (r *WatchlistsRepository) saveAssets(watchlist *Watchlist) error {
	_, err := r.db.Delete(ASSETS_TABLE_NAME, squirrel.Eq{"watchlist_id": watchlist.Id})

	if err != nil {
		return err
	}

	for _, asset := range watchlist.assets {
		configs, err := json.Marshal(asset.Settings)

		_, err = r.db.Insert(ASSETS_TABLE_NAME, map[string]interface{}{
			"watchlist_id": watchlist.Id,
			"ticker_code":  asset.TickerCode,
			"ticker_type":  asset.Kind,
			"configs":      configs,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
