CREATE TABLE IF NOT EXISTS users(
  id text NOT NULL PRIMARY KEY,
  name varchar(240) NOT NULL,
	email varchar(240) NOT NULL UNIQUE,
	user_pwd text NOT NULL,
	phone_number varchar(18) NOT NULL UNIQUE,
	created_at timestamp NOT NULL DEFAULT NOW(),
	updated_at timestamp NOT NULL DEFAULT NOW(),
	deleted_at timestamp NULL DEFAULT NULL
);