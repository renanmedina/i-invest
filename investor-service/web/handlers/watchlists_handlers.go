package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/renanmedina/i-invest/internal/wallet_monitoring"
	"github.com/renanmedina/i-invest/internal/wallet_monitoring/watchlists"
)

func ImportWatchlistFromB3SummaryReport(c *gin.Context) {
	_, reportFilePath := ImportB3FileFormHandler(c)
	use_case := watchlists.NewImportWatchlistFromB3Summary()
	id, _ := uuid.Parse("db1957f7-3d9e-4c4d-8e6d-72e7f40d6803")
	watchlist, err := use_case.Execute(id, reportFilePath, nil)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"watchlist": watchlist,
	})
}

func FetchNewAnnouncements(c *gin.Context) {
	use_case := wallet_monitoring.NewFetchAnnouncementsForWatchedTickers()
	use_case.Execute()
	c.JSON(http.StatusOK, nil)
}
