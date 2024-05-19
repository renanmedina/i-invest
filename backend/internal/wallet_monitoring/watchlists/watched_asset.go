package watchlists

type WatchedAsset struct {
	TickerCode string              `json:"ticker_code"`
	Kind       string              `json:"ticker_type"`
	Settings   WatchedAssetSetting `json:"settings"`
	Persisted  bool
}

type WatchedAssetSetting struct {
	MonitorAnnouncementsEnabled           bool `json:"monitor_announcements_enabled"`
	MonitorOccupationRateVariationEnabled bool `json:"monitor_occupation_rate_variation_enabled"`
}

func NewWatchedAsset(tickerCode string, TickerType string) WatchedAsset {
	return WatchedAsset{
		TickerCode: tickerCode,
		Kind:       TickerType,
		Settings:   WatchedAssetSetting{true, true},
	}
}
