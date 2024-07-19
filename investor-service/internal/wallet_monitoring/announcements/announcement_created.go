package announcements

const COMPANY_ANNOUNCEMENT_CREATED_EVENT_NAME = "CompanyAnnouncementCreated"

type CompanyAnnouncementCreated struct {
	announcement *CompanyAnnouncement
}

func NewCompanyAnnouncementCreatedEvent(announcement *CompanyAnnouncement) *CompanyAnnouncementCreated {
	return &CompanyAnnouncementCreated{announcement}
}

func (e *CompanyAnnouncementCreated) Name() string {
	return COMPANY_ANNOUNCEMENT_CREATED_EVENT_NAME
}

func (e *CompanyAnnouncementCreated) ObjectId() string {
	return e.announcement.Id.String()
}

func (e *CompanyAnnouncementCreated) ObjectType() string {
	return "CompanyAnnouncement"
}

func (e *CompanyAnnouncementCreated) EventData() map[string]interface{} {
	return map[string]interface{}{
		"announcement": e.announcement,
	}
}
