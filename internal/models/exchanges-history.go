// Package models contains application structs
package models

type ExchangesHistory struct {
	Date           string
	Cryptoamount   float64
	Fiatamount     float64
	Fee            float64
	Cryptocurrency string
	Paymethod      string
	Type           string
	Status         string
}
