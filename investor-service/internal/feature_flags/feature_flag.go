package feature_flags

type FeatureFlag struct {
	FeatureName string        `json:"feature_name"`
	Enabled     bool          `json:"enabled"`
	Settings    []interface{} `json:"settings"`
	persisted   bool
}

func NewFeatureFlag(name string, enabled bool, settings []interface{}) FeatureFlag {
	return FeatureFlag{
		FeatureName: name,
		Enabled:     enabled,
		Settings:    settings,
	}
}

func (f *FeatureFlag) Persisted() bool {
	return f.persisted
}

func (f *FeatureFlag) Feature() string {
	return f.FeatureName
}
