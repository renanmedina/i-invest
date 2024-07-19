package accounts

import (
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/renanmedina/i-invest/internal/wallet_monitoring/watchlists"
	"github.com/renanmedina/i-invest/utils"
)

type UsersRepository struct {
	db *utils.DatabaseAdapdater
}

const USERS_TABLE_NAME = "users"

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{utils.GetDatabase()}
}

func (r *UsersRepository) GetWithAnnoucementMonitoringEnabledByTickerCode(tickerCode string) (User, error) {
	query := squirrel.
		Select("DISTINCT users.id, users.name, users.email, users.phone_number").
		Join(fmt.Sprintf("%s ON %s.id = %s.user_id", watchlists.TABLE_NAME, USERS_TABLE_NAME, watchlists.TABLE_NAME)).
		Join(fmt.Sprintf("%s ON %s.watchlist_id = %s.id", watchlists.ASSETS_TABLE_NAME, watchlists.ASSETS_TABLE_NAME, watchlists.TABLE_NAME)).
		From(USERS_TABLE_NAME).
		PlaceholderFormat(squirrel.Dollar).
		Where(fmt.Sprintf("%s.ticker_code = ?", watchlists.ASSETS_TABLE_NAME), tickerCode).
		RunWith(r.db.GetConnection())

	user, err := buildUserFromDb(query.QueryRow())

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func buildUserFromDb(dbRow squirrel.RowScanner) (User, error) {
	var user User
	dbRow.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.PhoneNumber,
	)

	if user.Id.String() == "" {
		return User{}, errors.New("can't find User")
	}

	user.persisted = true
	return user, nil
}
