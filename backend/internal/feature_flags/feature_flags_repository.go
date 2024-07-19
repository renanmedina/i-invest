package feature_flags

import (
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/renanmedina/i-invest/utils"
)

type FeatureFlagsRepository struct {
	db *utils.DatabaseAdapdater
}

const FLAGS_TABLE_NAME = "feature_flags"

func NewFeatureFlagsRepository() FeatureFlagsRepository {
	return FeatureFlagsRepository{
		utils.GetDatabase(),
	}
}

func (r *FeatureFlagsRepository) GetByFlagName(name string) (*FeatureFlag, error) {
	query := squirrel.
		Select("*").
		From(FLAGS_TABLE_NAME).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"feature_name": name}).
		RunWith(r.db.GetConnection())

	flag, err := buildFromDb(query.QueryRow())

	if err != nil {
		return nil, err
	}

	return &flag, nil
}

func (r *FeatureFlagsRepository) Save(flag FeatureFlag) (*FeatureFlag, error) {

	return nil, nil
}

func buildFromDb(dbRow squirrel.RowScanner) (FeatureFlag, error) {
	var featureFlag FeatureFlag
	dbRow.Scan(
		&featureFlag.FeatureName,
		&featureFlag.Enabled,
		&featureFlag.Settings,
	)

	if featureFlag.FeatureName == "" {
		return FeatureFlag{}, errors.New("FeatureFlag not found")
	}

	featureFlag.persisted = true
	return featureFlag, nil
}
