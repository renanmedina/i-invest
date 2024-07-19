package announcements

const COMPANY_ANNOUNCEMENT_TRANSCRIBED_SUCCESSFULLY_EVENT_NAME = "CompanyAnnouncementTranscribedSuccessfully"

type CompanyAnnouncementTranscribedSuccessfully struct {
	announcement *CompanyAnnouncement
}

func NewCompanyAnnouncementTranscribedSuccessfully(announcement *CompanyAnnouncement) *CompanyAnnouncementTranscribedSuccessfully {
	return &CompanyAnnouncementTranscribedSuccessfully{announcement}
}

func (e *CompanyAnnouncementTranscribedSuccessfully) Name() string {
	return COMPANY_ANNOUNCEMENT_TRANSCRIBED_SUCCESSFULLY_EVENT_NAME
}

func (e *CompanyAnnouncementTranscribedSuccessfully) ObjectId() string {
	return e.announcement.Id.String()
}

func (e *CompanyAnnouncementTranscribedSuccessfully) ObjectType() string {
	return "CompanyAnnouncement"
}

func (e *CompanyAnnouncementTranscribedSuccessfully) EventData() map[string]interface{} {
	return map[string]interface{}{
		"announcement": e.announcement,
	}
}
