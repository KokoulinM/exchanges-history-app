# Postgres

Run server:

````
pg_ctl -D /opt/homebrew/var/postgres start
````

Start psql and open database postgres, which is the database postgres uses itself to store roles, permissions, and structure:
````
psql postgres
````

````
CREATE SCHEMA IF NOT EXISTS exchanges_history;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS history (
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    date TIMESTAMP NOT NULL,
    crypto_amount FLOAT NOT NULL,
    fiat_amount FLOAT NOT NULL,
    fee FLOAT NOT NULL,
    crypto_currency VARCHAR NOT NULL,
    pay_method VARCHAR NOT NULL,
    type VARCHAR NOT NULL,
    status VARCHAR NOT NULL
);
````
# Test

````
Check coverage
go test ./... -v -short -p 1 -cover
go test ./...
````
# Linting

````
golangci-lint run
````

# Frontend
````
cd ./web 

yarn/npm i - install npm packages

dev - yarn start
build - yarn build
````