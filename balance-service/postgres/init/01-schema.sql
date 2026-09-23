CREATE TABLE IF NOT EXISTS balances (
    account_id UUID PRIMARY KEY,
    balance NUMERIC(15,2) NOT NULL DEFAULT 0,
    updated_at TIMESTAMP NOT NULL
);
