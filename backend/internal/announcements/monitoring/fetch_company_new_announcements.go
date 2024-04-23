package announcements

import (
	"github.com/renanmedina/investment-warlock/internal/market/status_invest"
)

type FetchCompanyNewAnnouncement struct {
	announcementsService *status_invest.AnnouncementService
}

func NewFetchTickerNewAnnouncements() *FetchCompanyNewAnnouncement {
	return &FetchCompanyNewAnnouncement{
		announcementsService: status_invest.NewAnnouncementsService(),
	}
}

func (uc *FetchCompanyNewAnnouncement) execute(tickerCode string, year int) bool {
	_, err := uc.announcementsService.GetByTickerCodeAndYear(tickerCode, year)

	if err != nil {
		panic(err.Error())
	}

	return true
}
