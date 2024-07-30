CREATE TABLE IF NOT EXISTS wallet_transactions(
  id text NOT NULL PRIMARY KEY,
  wallet_id text NOT NULL REFERENCES wallets,
  transaction_type varchar(5) NOT NULL,
  transaction_date date NOT NULL,
  asset_kind varchar(35) NOT NULL,
  asset_code varchar(200),
  asset_description text NULL,
  quantity decimal(10, 8) NOT NULL,
  unit_price decimal(10, 8) NOT NULL,
  total_amount decimal(10, 8) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL DEFAULT NULL
);