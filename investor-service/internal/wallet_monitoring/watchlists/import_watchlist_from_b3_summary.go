package watchlists

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/renanmedina/i-invest/internal/wallets/b3"
	"github.com/renanmedina/i-invest/utils"
)

type ImportWatchlistFromB3Summary struct {
	allWatchlists *WatchlistsRepository
	logger        *utils.ApplicationLogger
}

func NewImportWatchlistFromB3Summary() *ImportWatchlistFromB3Summary {
	return &ImportWatchlistFromB3Summary{
		NewWatchlistsRepository(),
		utils.GetApplicationLogger(),
	}
}

func (uc *ImportWatchlistFromB3Summary) Execute(userId uuid.UUID, filepath string, watchlistId *string) (*Watchlist, error) {
	uc.logger.Info(fmt.Sprintf("Importing watchlist from b3 for user %s summary report file at %s", userId, filepath))

	parsedItems, err := b3.ParseSummaryReport(filepath)

	if err != nil {
		uc.logger.Error(fmt.Sprintf("Failed to parse b3 summary at %s with: %s", filepath, err.Error()))
		return nil, err
	}

	watchlistObj := NewEmptyWatchlist(userId)

	if watchlistId != nil {
		existingWatchlist, err := uc.allWatchlists.GetById(*watchlistId)

		if err != nil {
			logMsg := fmt.Sprintf("Failed to import watchlist %s from b3 summary of %s to watchlist %s with: %s", filepath, *watchlistId, err.Error())
			uc.logger.Error(logMsg)
			return nil, err
		}

		watchlistObj = *existingWatchlist
	}

	watchlistObj.importFromB3(&parsedItems)
	uc.allWatchlists.Save(&watchlistObj)
	return &watchlistObj, nil
}
