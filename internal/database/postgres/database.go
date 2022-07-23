// Package postgres provides data storage in a postgres database
package postgres

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"
)

type PostgresDatabase struct {
	conn   *sql.DB
	logger *zerolog.Logger
}

func New(db *sql.DB, logger *zerolog.Logger) *PostgresDatabase {
	return &PostgresDatabase{
		conn:   db,
		logger: logger,
	}
}

func (db *PostgresDatabase) Ping(ctx context.Context) error {
	err := db.conn.PingContext(ctx)
	if err != nil {
		db.logger.Error().Msg(err.Error())
		return err
	}
	return nil
}
