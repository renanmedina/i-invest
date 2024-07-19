CREATE TABLE IF NOT EXISTS company_announcements(
  id text NOT NULL PRIMARY KEY,
	ticker_code varchar(8) NOT NULL,
	title text NOT NULL,
  announcement_type text NOT NULL,
	announcement_date timestamp NOT NULL,
	file_url text NULL DEFAULT NULL,
	original_file_url text NULL DEFAULT NULL,
	created_at timestamp NOT NULL DEFAULT NOW(),
	updated_at timestamp NOT NULL DEFAULT NOW(),
	deleted_at timestamp NULL DEFAULT NULL
);

CREATE INDEX idx_ticker_code_announcement ON company_announcements(ticker_code);
CREATE INDEX idx_type_announcement ON company_announcements(announcement_type);
CREATE INDEX idx_date_announcement ON company_announcements(announcement_date);