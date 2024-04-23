package event_store

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/renanmedina/investment-warlock/utils"
)

const TABLE_NAME = "events"

type EventsRepository struct {
	db *sql.DB
}

func NewEventsRepository() *EventsRepository {
	return &EventsRepository{
		db: utils.GetDatabase(),
	}
}

func (r *EventsRepository) Save(event PublishableEvent) error {
	eventData, err := json.Marshal(event.EventData())

	if err != nil {
		fmt.Println("Failed marshal event data")
		panic(err)
	}

	query := squirrel.
		Insert(TABLE_NAME).
		Columns("event_name", "object_id", "object_type", "event_data").
		Values(event.Name(), event.ObjectId(), event.ObjectType(), eventData).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	_, err = query.Exec()

	if err != nil {
		panic(err)
		return err
	}

	return nil
}
