package announcements

import (
	"database/sql"

	"github.com/renanmedina/investment-warlock/utils"
)

type AllAnnouncements struct {
	db *sql.DB
}

func NewAllAnnouncements() *AllAnnouncements {
	return &AllAnnouncements{
		db: utils.GetDatabase(),
	}
}

func (r *AllAnnouncements) GetByTickerCodeAndYear(tickerCode string, year int) {

}
