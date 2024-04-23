package announcements

import (
	"database/sql"
	"errors"
	"fmt"

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

func (r *AnnouncementsRepository) GetById(id string) (*CompanyAnnouncement, error) {
	query := squirrel.Select("*").From(TABLE_NAME).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": id}).
		Where("deleted_at is null").
		Limit(1).
		RunWith(r.db)

	row := query.QueryRow()

	var announcement CompanyAnnouncement
	row.Scan(
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
		return nil, errors.New(fmt.Sprintf("Can't find CompanyAnnouncement with ID: %s", id))
	}

	return &announcement, nil
}

func (r *AnnouncementsRepository) GetByTickerCodeAndYear(tickerCode string, year int) []CompanyAnnouncement {
	query := squirrel.Select("*").From(TABLE_NAME).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"ticker_code": tickerCode}).
		Where("DATE_PART('year', announcement_date::timestamp) = ?", year).
		Where("deleted_at is null").
		RunWith(r.db)

	rows, err := query.Query()
	list := make([]CompanyAnnouncement, 0)

	if err != nil {
		fmt.Println(query.ToSql())
		panic(err)
	}

	for rows.Next() {
		var announcement CompanyAnnouncement
		rows.Scan(
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
		list = append(list, announcement)
	}

	return list
}

func (r *AnnouncementsRepository) Save(announcement CompanyAnnouncement) error {
	query := squirrel.Insert(TABLE_NAME).
		Columns("id", "ticker_code", "title", "announcement_type",
			"announcement_date", "file_url", "original_file_url",
		).Values(
		announcement.Id, announcement.TickerCode, announcement.Subject,
		announcement.AnnouncementType, announcement.AnnouncementDate, announcement.FileUrl,
		announcement.OriginalFileUrl,
	).RunWith(r.db).PlaceholderFormat(squirrel.Dollar)

	_, err := query.Exec()

	if err != nil {
		return err
	}

	return nil
}
