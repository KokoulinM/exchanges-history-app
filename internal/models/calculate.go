package models

type ResponseCalculation struct {
	// Сумма сделок в руб
	FiatAmounts float64 `json:"fiatAmounts"`
	// сумма в BTC (за вычетом Fee)
	CryptoAmount float64 `json:"cryptoAmount"`
	// средняя стоимость крипты за вычетом комиссии
	AverageCost float64 `json:"averageCost"`
}
