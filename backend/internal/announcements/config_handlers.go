package announcements

import "github.com/renanmedina/investment-warlock/internal/event_store"

func configureEventHandlers() map[string][]event_store.EventHandler {
	return map[string][]event_store.EventHandler{
		COMPANY_ANNOUNCEMENT_CREATED_EVENT_NAME: {
			event_store.NewSaveEventToStoreHandler(),
			NewDownloadAndSaveAnnouncementFileHandler(),
		},
		COMPANY_ANNOUNCEMENT_FILE_DOWNLOADED_AND_SAVED_EVENT_NAME: {
			event_store.NewSaveEventToStoreHandler(),
		},
	}
}
