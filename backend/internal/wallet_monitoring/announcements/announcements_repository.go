package announcements

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/renanmedina/investment-warlock/utils"
)

const TABLE_NAME = "company_announcements"

type AnnouncementsRepository struct {
	db *sql.DB
}

func NewAnnouncementsRepository() *AnnouncementsRepository {
	return &AnnouncementsRepository{
		db: utils.GetDatabase(),
	}
}

func BuildAnnouncementFromDb(dbRow squirrel.RowScanner) (*CompanyAnnouncement, error) {
	var announcement CompanyAnnouncement
	dbRow.Scan(
		&announcement.Id,
		&announcement.TickerCode,
		&announcement.Subject,
		&announcement.AnnouncementType,
		&announcement.AnnouncementDate,
		&announcement.FileUrl,
		&announcement.OriginalFileUrl,
		&announcement.CreatedAt,
		&announcement.UpdatedAt,
		&announcement.DeletedAt,
	)

	if announcement.TickerCode == "" {
		return nil, errors.New("can't find CompanyAnnouncement")
	}

	announcement.Persisted = true
	return &announcement, nil
}

func BuildAnnouncementsListFromDb(dbRows *sql.Rows) []CompanyAnnouncement {
	list := make([]CompanyAnnouncement, 0)

	for dbRows.Next() {
		var announcement CompanyAnnouncement
		dbRows.Scan(
			&announcement.Id,
			&announcement.TickerCode,
			&announcement.Subject,
			&announcement.AnnouncementType,
			&announcement.AnnouncementDate,
			&announcement.FileUrl,
			&announcement.OriginalFileUrl,
			&announcement.CreatedAt,
			&announcement.UpdatedAt,
			&announcement.DeletedAt,
		)
		announcement.Persisted = true
		list = append(list, announcement)
	}

	return list
}

func (r *AnnouncementsRepository) GetById(id string) (*CompanyAnnouncement, error) {
	query := squirrel.Select("*").From(TABLE_NAME).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": id}).
		Where("deleted_at is null").
		Limit(1).
		RunWith(r.db)

	announcement, err := BuildAnnouncementFromDb(query.QueryRow())

	if err != nil {
		return nil, fmt.Errorf("can't find CompanyAnnouncement with ID: %s", id)
	}

	return announcement, nil
}

func (r *AnnouncementsRepository) GetByTickerCodeAndYear(tickerCode string, year int) []CompanyAnnouncement {
	query := squirrel.Select("*").From(TABLE_NAME).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"ticker_code": tickerCode}).
		Where("DATE_PART('year', announcement_date::timestamp) = ?", year).
		Where("deleted_at is null").
		RunWith(r.db)

	rows, err := query.Query()

	if err != nil {
		fmt.Println(query.ToSql())
		panic(err)
	}

	announcements := BuildAnnouncementsListFromDb(rows)
	return announcements
}

func (r *AnnouncementsRepository) Save(announcement *CompanyAnnouncement) (*CompanyAnnouncement, error) {
	var err error

	if announcement.Persisted {
		_, err = squirrel.Update(TABLE_NAME).
			Set("file_url", announcement.FileUrl).
			Set("updated_at", time.Now()).
			Where(squirrel.Eq{"id": announcement.Id}).
			RunWith(r.db).
			PlaceholderFormat(squirrel.Dollar).
			Exec()
	} else {
		_, err = squirrel.Insert(TABLE_NAME).
			Columns("id", "ticker_code", "title", "announcement_type",
				"announcement_date", "file_url", "original_file_url",
			).Values(
			announcement.Id, announcement.TickerCode, announcement.Subject,
			announcement.AnnouncementType, announcement.AnnouncementDate, announcement.FileUrl,
			announcement.OriginalFileUrl,
		).
			RunWith(r.db).
			PlaceholderFormat(squirrel.Dollar).
			Exec()
	}

	if err != nil {
		return nil, err
	}

	return announcement, nil
}
