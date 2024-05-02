CREATE TABLE IF NOT EXISTS watchlists(
  id text NOT NULL PRIMARY KEY,
  user_id text NOT NULL REFERENCES users,
  label varchar(200) NOT NULL,
	created_at timestamp NOT NULL DEFAULT NOW(),
	updated_at timestamp NOT NULL DEFAULT NOW(),
	deleted_at timestamp NULL DEFAULT NULL
);