CREATE TABLE IF NOT EXISTS watchlists_assets_configs(
  id SERIAL NOT NULL PRIMARY KEY,
  watchlists_asset_id bigint REFERENCES watchlists_assets,
  ticker_code varchar(8) NOT NULL,
  configs jsonb NOT NULL,
	created_at timestamp NOT NULL DEFAULT NOW(),
	updated_at timestamp NOT NULL DEFAULT NOW(),
	deleted_at timestamp NULL DEFAULT NULL
);
