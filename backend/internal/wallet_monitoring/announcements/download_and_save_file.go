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
	logger           *utils.ApplicationLogger
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
		utils.GetApplicationLogger(),
	}
}

func (uc *DownloadAndSaveAnnouncementFile) Execute(announcementId string) {
	// rand.Seed(time.Now().UnixNano())
	// n := rand.Intn(25) // n will be between 0 and 10
	// time.Sleep(time.Duration(n) * time.Second)
	announcement, errFind := uc.allAnnouncements.GetById(announcementId)

	uc.logger.Info(fmt.Sprintf("Downloading file for announcement: %s", announcementId))

	if errFind != nil {
		uc.logger.Info(fmt.Sprintf("Failed to download announcement file with: %s not found", announcementId))
		return
	}

	if announcement.FileUrl != "" {
		return
	}

	tmpFilePath, errDownload := downloadFileToTmpDir(uc.logger, *announcement)

	if errDownload != nil {
		uc.logger.Error(fmt.Sprintf("Failed to download announcement file for %s with: %s", announcementId, errDownload.Error()))
		// retry in a go-routine
		uc.logger.Info(fmt.Sprintf("Retrying download file for announcement: %s", announcementId))
		go uc.Execute(announcementId)
		return
	}

	uploadPath := announcement.MakeUploadPath()
	logMsg := fmt.Sprintf("Uploading announcement file for %s from %s to %s", announcementId, announcement.OriginalFileUrl, uploadPath)
	uc.logger.Info(logMsg)

	storageFileUrl, errUpload := uc.fileStorage.Upload(tmpFilePath, announcement.MakeUploadPath())

	if errUpload != nil {
		uc.logger.Error(fmt.Sprintf("Failed to upload file for %s from %s to %s with: %s", announcementId, uploadPath, errUpload.Error()))
		return
	}

	announcement.FileUrl = storageFileUrl
	uc.allAnnouncements.Save(announcement)
	uc.logger.Info(fmt.Sprintf("Announcement file for %s downloaded and uploaded successfuly at %s", announcementId, announcement.FileUrl))
	event := NewCompanyAnnouncementFileDownloadedAndSaved(announcement)
	uc.eventPublisher.Publish(event)
}

func downloadFileToTmpDir(logger *utils.ApplicationLogger, announcement CompanyAnnouncement) (string, error) {
	// Get the data
	fileResponse, err := http.Get(announcement.OriginalFileUrl)
	if err != nil {
		return "", err
	}

	defer fileResponse.Body.Close()

	// Check server response
	if fileResponse.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", fileResponse.Status)
	}

	// Writer the body to file
	tmpFilePath := fmt.Sprintf("/tmp/%s", announcement.MakeTempFileName())
	logMsg := fmt.Sprintf("Downloading announcement file for %s from %s to %s", announcement.Id, announcement.OriginalFileUrl, tmpFilePath)
	logger.Info(logMsg)

	remoteFileContent, readFileErr := io.ReadAll(fileResponse.Body)

	if readFileErr != nil {
		return "", readFileErr
	}

	errWriteFile := os.WriteFile(tmpFilePath, remoteFileContent, os.ModePerm)
	if errWriteFile != nil {
		return "", errWriteFile
	}

	return tmpFilePath, nil
}
