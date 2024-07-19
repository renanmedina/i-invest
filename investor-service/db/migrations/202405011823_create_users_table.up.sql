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

-- INSERT INTO users(id, name, email, user_pwd, phone_number) 
-- VALUES('db1957f7-3d9e-4c4d-8e6d-72e7f40d6803', 'Renan Medina', 'renan@silvamedina.com.br', '123456', '+5521979582480');