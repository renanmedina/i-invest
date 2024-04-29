package announcements

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/renanmedina/investment-warlock/internal/event_store"
	"github.com/renanmedina/investment-warlock/utils"
)

type DownloadAndSaveAnnouncementFileHandler struct{}

func (h *DownloadAndSaveAnnouncementFileHandler) Handle(event event_store.PublishableEvent) {
	use_case := NewDownloadAndSaveAnnouncementFile()
	use_case.Execute(event.ObjectId())
}

func NewDownloadAndSaveAnnouncementFileHandler() *DownloadAndSaveAnnouncementFileHandler {
	return &DownloadAndSaveAnnouncementFileHandler{}
}

type DownloadAndSaveAnnouncementFile struct {
	allAnnouncements *AnnouncementsRepository
	fileStorage      utils.FileStorageService
	eventPublisher   *event_store.EventPublisher
}

func NewDownloadAndSaveAnnouncementFile() *DownloadAndSaveAnnouncementFile {
	fileService, err := utils.NewS3FileStorage()

	if err != nil {
		panic(err)
	}

	return &DownloadAndSaveAnnouncementFile{
		NewAnnouncementsRepository(),
		fileService,
		event_store.NewEventPublisherWith(configureEventHandlers()),
	}
}

func (uc *DownloadAndSaveAnnouncementFile) Execute(announcementId string) {
	announcement, err := uc.allAnnouncements.GetById(announcementId)

	if err != nil {
		panic(err)
	}

	if announcement.FileUrl != "" {
		return
	}

	tmpFilePath := downloadFileToTmpDir(*announcement)
	storageFileUrl, err := uc.fileStorage.Upload(tmpFilePath, announcement.MakeUploadPath())

	if err != nil {
		panic(err)
	}

	announcement.FileUrl = storageFileUrl
	uc.allAnnouncements.Save(announcement)
	event := NewCompanyAnnouncementFileDownloadedAndSaved(announcement)
	uc.eventPublisher.Publish(event)
}

func downloadFileToTmpDir(announcement CompanyAnnouncement) string {
	// Get the data
	fileResponse, err := http.Get(announcement.OriginalFileUrl)
	if err != nil {
		panic(err)
	}

	defer fileResponse.Body.Close()

	// Check server response
	if fileResponse.StatusCode != http.StatusOK {
		panic(fmt.Errorf("bad status: %s", fileResponse.Status))
	}

	// Writer the body to file
	tmpFilePath := fmt.Sprintf("/tmp/%s", announcement.MakeTempFileName())
	remoteFileContent, err := io.ReadAll(fileResponse.Body)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile(tmpFilePath, remoteFileContent, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return tmpFilePath
}
