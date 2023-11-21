CREATE SCHEMA sso;
CREATE TABLE IF NOT EXISTS sso.accounts (
    id SERIAL PRIMARY KEY,
    login TEXT UNIQUE,
    password TEXT
);
CREATE TABLE IF NOT EXISTS sso.tokens (
    id SERIAL PRIMARY KEY,
    token_value TEXT UNIQUE
);
CREATE TABLE IF NOT EXISTS sso.accounts_tokens (
    id SERIAL PRIMARY KEY,
    account_id int,
    token_id int,
    CONSTRAINT fk_account FOREIGN KEY(account_id) REFERENCES sso.accounts(id),
    CONSTRAINT fk_token FOREIGN KEY(token_id) REFERENCES sso.tokens(id)
);