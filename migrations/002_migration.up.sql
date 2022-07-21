CREATE SCHEMA IF NOT EXISTS exchanges_history;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS history (
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    data VARCHAR NOT NULL,
    crypto_amount FLOAT NOT NULL,
    fiat_amount FLOAT NOT NULL,
    fee FLOAT NOT NULL,
    crypto_currency VARCHAR NOT NULL,
    pay_method VARCHAR NOT NULL,
    type VARCHAR NOT NULL,
    status VARCHAR NOT NULL
);