// Package models contains application structs
package models

import (
	"time"
)

type ExchangesHistory struct {
	Date           time.Time `json:"date"`
	Cryptoamount   float64   `json:"crypto_amount"`
	Fiatamount     float64   `json:"fiat_amount"`
	Fee            float64   `json:"fee"`
	Cryptocurrency string    `json:"crypto_currency"`
	Paymethod      string    `json:"pay_method"`
	Type           string    `json:"type"`
	Status         string    `json:"status"`
}

func (eh ExchangesHistory) MarshalJSON() string {
	return eh.Date.Format(time.RFC3339)
}

type ResponseExchangesHistoryInfo struct {
	CryptoCurrencies []string `json:"cryptoCurrencies"`
	PayMethods       []string `json:"payMethods"`
}
