// Package csv contains parsing methods
package csv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/KokoulinM/exchanges-history-app/internal/models"
)

const (
	DateCol           int = 0
	CryptoamountCol   int = 2
	FiatamountCol     int = 3
	FeeCol            int = 4
	CryptocurrencyCol int = 7
	PaymethodCol      int = 8
	TypeCol           int = 11
	StatusCol         int = 12
)

func parser(data [][]string, filters ...func(eh models.ExchangesHistory) bool) ([]models.ExchangesHistory, error) {
	var exchangesHistory []models.ExchangesHistory

	for i, line := range data {
		if i > 0 {
			var rec models.ExchangesHistory
			for col, field := range line {
				trimmed := strings.TrimSpace(field)
				switch col {
				case DateCol:
					rec.Date = trimmed
				case CryptoamountCol:
					fl, err := strconv.ParseFloat(trimmed, 32)
					if err != nil {
						return nil, err
					}
					rec.Cryptoamount = fl
				case FiatamountCol:
					fl, err := strconv.ParseFloat(trimmed, 32)
					if err != nil {
						return nil, err
					}
					rec.Fiatamount = fl
				case FeeCol:
					fl, err := strconv.ParseFloat(trimmed, 32)
					if err != nil {
						return nil, err
					}
					rec.Fee = fl
				case CryptocurrencyCol:
					rec.Cryptocurrency = trimmed
				case PaymethodCol:
					rec.Paymethod = trimmed
				case TypeCol:
					rec.Type = trimmed
				case StatusCol:
					rec.Status = trimmed
				default:
					log.Printf("The column %d does not fit the conditions of the function", col)
				}
			}

			if combineFilters(rec, filters...) {
				exchangesHistory = append(exchangesHistory, rec)
			}
		}
	}

	return exchangesHistory, nil
}

func Reader(name string) ([]models.ExchangesHistory, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	exchangesHistory, err := parser(data, filterType, filterStatus)
	if err != nil {
		return nil, err
	}

	return exchangesHistory, nil
}
