package announcements

import (
	"fmt"

	"github.com/renanmedina/investment-warlock/internal/event_store"
)

type TranscribeAnnouncementFileHandler struct{}

func NewTranscribeAnnouncementFileHandler() *TranscribeAnnouncementFileHandler {
	return &TranscribeAnnouncementFileHandler{}
}

func (h *TranscribeAnnouncementFileHandler) Handle(event event_store.PublishableEvent) {
	uc := NewTranscribeAnnouncementFile()
	uc.Execute(event.ObjectId())
}

type TranscribeAnnouncementFile struct {
	allAnnouncements *AnnouncementsRepository
	ocrService       interface{}
	audioService     interface{}
	eventPublisher   *event_store.EventPublisher
}

func NewTranscribeAnnouncementFile() *TranscribeAnnouncementFile {
	return &TranscribeAnnouncementFile{
		NewAnnouncementsRepository(),
		nil,
		nil,
		event_store.NewEventPublisherWith(configureEventHandlers()),
	}
}

func (uc *TranscribeAnnouncementFile) Execute(announcementId string) {
	announcement, err := uc.allAnnouncements.GetById(announcementId)

	if err != nil {
		panic(err)
	}

	fmt.Println(announcement)

	// transcriptionText, err := uc.ocrService.GetTextFromFile(announcement.FileUrl)

	// if err != nil {
	// 	panic(err)
	// }

	// audioFileUrl, err := uc.audioService.CreateFromText(transcriptionText)

	// if err != nil {
	// 	panic(err)
	// }

	// announcement.addTranscription(transcriptionText, audioFileUrl)
	// uc.allAnnouncements.Save(announcement)
	// uc.eventPublisher.Publish(NewCompanyAnnouncementTranscribedSuccessfully(announcement))
}
