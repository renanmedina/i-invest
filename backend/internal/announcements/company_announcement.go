package announcements

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/renanmedina/investment-warlock/internal/market"
)

type CompanyAnnouncement struct {
	Id               uuid.UUID `json:"id"`
	TickerCode       string    `json:"ticker_code"`
	Subject          string    `json:"title"`
	AnnouncementType string    `json:"type"`
	FileUrl          string    `json:"file_url"`
	OriginalFileUrl  string    `json:"original_file_url"`
	AnnouncementDate time.Time `json:"date"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        time.Time `json:"deleted_at"`
}

func (a CompanyAnnouncement) MakeTempFileName() string {
	return fmt.Sprintf("announcement_file_%s_%s.pdf", a.TickerCode, a.Id.String())
}

func (a CompanyAnnouncement) MakeUploadPath() string {
	return fmt.Sprintf("%s/%s", a.TickerCode, a.MakeTempFileName())
}

func translateAnnouncementFromApiResults(tickerCode string, apiItems []market.AnnouncementApiItem) []CompanyAnnouncement {
	translateds := make([]CompanyAnnouncement, 0)

	for _, item := range apiItems {
		translateds = append(translateds, transalteAnnoutranslateAnnouncementFromApiItem(tickerCode, item))
	}

	return translateds
}

func transalteAnnoutranslateAnnouncementFromApiItem(tickerCode string, item market.AnnouncementApiItem) CompanyAnnouncement {
	date, _ := time.Parse("2006-01-02T00:00:00", item.ReferenceDate)

	return CompanyAnnouncement{
		Id:               uuid.New(),
		TickerCode:       tickerCode,
		Subject:          item.Subject,
		AnnouncementType: item.AnnouncementType,
		FileUrl:          "",
		OriginalFileUrl:  item.LinkUrl,
		AnnouncementDate: date,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}
