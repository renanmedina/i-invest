package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renanmedina/investment-warlock/internal/announcements"
	"github.com/renanmedina/investment-warlock/internal/market"
)

func MarketTickerAnnouncements(c *gin.Context) {
	tickerId, wasFound := c.Params.Get("tickerId")

	if !wasFound {
		c.JSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}

	service := market.NewAnnouncementsService()
	announcements, err := service.GetByTickerCodeAndYear(tickerId, 2024)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"announcements": announcements,
	})
}

func FetchCompanyNewAnnouncements(c *gin.Context) {
	tickerId, wasFound := c.Params.Get("tickerId")

	if !wasFound {
		c.JSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}

	uc := announcements.NewFetchCompanyNewAnnouncements()
	announcements := uc.Execute(tickerId, 2024)

	c.JSON(http.StatusOK, gin.H{
		"new_announcements": announcements,
	})
}
