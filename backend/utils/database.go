package utils

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type DatabaseAdapdater struct {
	db *sql.DB
}

type DbRecordable interface {
	Persisted() bool
}

var dbAdapter *DatabaseAdapdater

func init() {
	initDB()
}

func initDB() {
	configs := GetConfigs()
	openedDb, err := sql.Open("postgres", configs.DB_URI)

	if err != nil {
		panic(err)
	}

	dbAdapter = &DatabaseAdapdater{openedDb}
}

func GetDatabase() *DatabaseAdapdater {
	return dbAdapter
}

func (adapter *DatabaseAdapdater) GetConnection() *sql.DB {
	return adapter.db
}

func (adapter *DatabaseAdapdater) Insert(tableName string, fieldsAndValues map[string]interface{}) (bool, error) {
	columns := getKeys(fieldsAndValues)
	fieldValues := getValuesOrdered(columns, fieldsAndValues)

	_, errInsert := squirrel.
		Insert(tableName).
		Columns(columns...).
		Values(fieldValues...).
		RunWith(adapter.db).
		PlaceholderFormat(squirrel.Dollar).
		Exec()

	if errInsert != nil {
		return false, errInsert
	}

	return true, nil
}

func (adapter *DatabaseAdapdater) Update(tableName string, fieldsAndValues map[string]interface{}, wheres interface{}) (bool, error) {
	updateBuilder := squirrel.Update(tableName)

	for fieldName, fieldVal := range fieldsAndValues {
		updateBuilder = updateBuilder.Set(fieldName, fieldVal)
	}

	updateBuilder.Where(wheres)

	_, errUpdate := updateBuilder.RunWith(adapter.db).PlaceholderFormat(squirrel.Dollar).Exec()

	if errUpdate != nil {
		return false, errUpdate
	}

	return true, nil
}

func (adapter *DatabaseAdapdater) Delete(tableName string, wheres interface{}) (bool, error) {
	deleteBuilder := squirrel.Delete(tableName)
	deleteBuilder.Where(wheres)

	_, errDelete := deleteBuilder.RunWith(adapter.db).PlaceholderFormat(squirrel.Dollar).Exec()

	if errDelete != nil {
		return false, errDelete
	}

	return true, nil
}

func getKeys(mapVar map[string]interface{}) []string {
	keys := make([]string, len(mapVar))

	i := 0
	for key := range mapVar {
		keys[i] = key
		i++
	}

	return keys
}

func getValuesOrdered(columns []string, mapVar map[string]interface{}) []interface{} {
	vals := make([]interface{}, len(mapVar))

	i := 0
	for _, column := range columns {
		vals[i] = mapVar[column]
		i++
	}

	return vals
}
