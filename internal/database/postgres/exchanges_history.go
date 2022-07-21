package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/KokoulinM/exchanges-history-app/internal/models"
)

func (db *PostgresDatabase) UploadFile(ctx context.Context, exchangesHistory []models.ExchangesHistory) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO history (data, crypto_amount, fiat_amount, fee, crypto_currency, pay_method, type, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)
	if err != nil {
		return err
	}

	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	for _, eh := range exchangesHistory {
		if _, err = stmt.ExecContext(ctx, eh.Date, eh.Cryptoamount, eh.Fiatamount, eh.Fee, eh.Cryptocurrency, eh.Paymethod, eh.Type, eh.Status); err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
