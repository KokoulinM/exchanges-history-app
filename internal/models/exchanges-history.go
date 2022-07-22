// Package models contains application structs
package models

type ExchangesHistory struct {
	Date           string  `json:"date"`
	Cryptoamount   float64 `json:"crypto_amount"`
	Fiatamount     float64 `json:"fiat_amount"`
	Fee            float64 `json:"fee"`
	Cryptocurrency string  `json:"crypto_currency"`
	Paymethod      string  `json:"pay_method"`
	Type           string  `json:"type"`
	Status         string  `json:"status"`
}
