CREATE TABLE IF NOT EXISTS watchlists_assets(
  id SERIAL NOT NULL PRIMARY KEY,
  watchlist_id TEXT NOT NULL REFERENCES watchlists(id) ON DELETE CASCADE,
  ticker_code varchar(8) NOT NULL,
  ticker_type varchar(14) NOT NULL,
	created_at timestamp NOT NULL DEFAULT NOW(),
	updated_at timestamp NOT NULL DEFAULT NOW(),
	deleted_at timestamp NULL DEFAULT NULL
);
