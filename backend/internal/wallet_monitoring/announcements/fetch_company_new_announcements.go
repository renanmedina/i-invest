package announcements

import (
	"github.com/renanmedina/i-invest/internal/event_store"
	"github.com/renanmedina/i-invest/internal/market"
)

type FetchCompanyNewAnnouncement struct {
	allAnnouncements     *AnnouncementsRepository
	announcementsService *market.AnnouncementService
	eventPublisher       *event_store.EventPublisher
}

func NewFetchCompanyNewAnnouncements() *FetchCompanyNewAnnouncement {
	return &FetchCompanyNewAnnouncement{
		allAnnouncements:     NewAnnouncementsRepository(),
		announcementsService: market.NewAnnouncementsService(),
		eventPublisher:       event_store.NewEventPublisherWith(configureEventHandlers()),
	}
}

func (uc *FetchCompanyNewAnnouncement) Execute(tickerCode string, year int) []CompanyAnnouncement {
	result, err := uc.announcementsService.GetByTickerCodeAndYear(tickerCode, year)

	if err != nil {
		panic(err.Error())
	}

	announcements := result.Items
	if announcements == nil || (announcements != nil && len(announcements) == 0) {
		return make([]CompanyAnnouncement, 0)
	}

	translatedAnnouncements := translateAnnouncementFromApiResults(tickerCode, announcements)
	savedAnnouncements := uc.allAnnouncements.GetByTickerCodeAndYear(tickerCode, year)
	newAnnouncements := diffCompanyAnnouncements(
		&translatedAnnouncements,
		&savedAnnouncements,
	)

	for _, newAnnouncement := range newAnnouncements {
		uc.allAnnouncements.Save(&newAnnouncement)
		event := NewCompanyAnnouncementCreatedEvent(&newAnnouncement)
		uc.eventPublisher.Publish(event)
	}

	return newAnnouncements
}

func diffCompanyAnnouncements(pivotList *[]CompanyAnnouncement, savedList *[]CompanyAnnouncement) []CompanyAnnouncement {
	missingAnnouncements := make([]CompanyAnnouncement, 0)
	exitingFilesUrls := make(map[string]bool, 0)

	for _, savedItem := range *savedList {
		exitingFilesUrls[savedItem.OriginalFileUrl] = true
	}

	for _, announcement := range *pivotList {
		_, announcementExists := exitingFilesUrls[announcement.OriginalFileUrl]
		if !announcementExists {
			missingAnnouncements = append(missingAnnouncements, announcement)
		}
	}

	return missingAnnouncements
}
