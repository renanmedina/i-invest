package announcements

const COMPANY_ANNOUNCEMENT_FILE_DOWNLOADED_AND_SAVED_EVENT_NAME = "CompanyAnnouncementFileDownloadedAndSaved"

type CompanyAnnouncementFileDownloadedAndSaved struct {
	announcement *CompanyAnnouncement
}

func NewCompanyAnnouncementFileDownloadedAndSaved(announcement *CompanyAnnouncement) *CompanyAnnouncementFileDownloadedAndSaved {
	return &CompanyAnnouncementFileDownloadedAndSaved{announcement}
}

func (e *CompanyAnnouncementFileDownloadedAndSaved) Name() string {
	return COMPANY_ANNOUNCEMENT_FILE_DOWNLOADED_AND_SAVED_EVENT_NAME
}

func (e *CompanyAnnouncementFileDownloadedAndSaved) ObjectId() string {
	return e.announcement.Id.String()
}

func (e *CompanyAnnouncementFileDownloadedAndSaved) ObjectType() string {
	return "CompanyAnnouncement"
}

func (e *CompanyAnnouncementFileDownloadedAndSaved) EventData() map[string]interface{} {
	return map[string]interface{}{
		"announcement": e.announcement,
	}
}
