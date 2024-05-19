package watchlists

type WatchedAsset struct {
	TickerCode string              `json:"ticker_code"`
	Kind       string              `json:"ticker_type"`
	Settings   WatchedAssetSetting `json:"settings"`
	Persisted  bool
}

type WatchedAssetSetting struct {
	NotifyNewAnnouncements        bool `json:"notify_new_announcements"`
	NotifyOccupationRateVariation bool `json:"notify_occupation_rate_variation"`
}

func NewWatchedAsset(tickerCode string, TickerType string) WatchedAsset {
	return WatchedAsset{
		TickerCode: tickerCode,
		Kind:       TickerType,
		Settings:   WatchedAssetSetting{true, true},
	}
}
