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

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO history (date, crypto_amount, fiat_amount, fee, crypto_currency, pay_method, type, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)
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

func (db *PostgresDatabase) GetHistory(ctx context.Context) ([]models.ExchangesHistory, error) {
	var result []models.ExchangesHistory

	sqlGetHistory := `SELECT date, crypto_amount, fiat_amount, fee, crypto_currency, pay_method, type, status FROM history;`
	rows, err := db.conn.QueryContext(ctx, sqlGetHistory)
	if err != nil {
		return result, err
	}
	if rows.Err() != nil {
		return result, rows.Err()
	}
	defer rows.Close()

	for rows.Next() {
		var eh models.ExchangesHistory
		err = rows.Scan(&eh.Date, &eh.Cryptoamount, &eh.Fiatamount, &eh.Fee, &eh.Cryptocurrency, &eh.Paymethod, &eh.Type, &eh.Status)
		if err != nil {
			return result, err
		}
		result = append(result, eh)
	}

	return result, nil
}
