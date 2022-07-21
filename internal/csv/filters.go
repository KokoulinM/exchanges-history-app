package csv

import "github.com/KokoulinM/exchanges-history-app/internal/models"

// нужны только purchase
func filterType(eh models.ExchangesHistory) bool {
	return eh.Type == "purchase"
}

// парсить только confirm_payment
func filterStatus(eh models.ExchangesHistory) bool {
	return eh.Status == "confirm_payment"
}

func combineFilters(eh models.ExchangesHistory, filters ...func(eh models.ExchangesHistory) bool) bool {
	if len(filters) == 0 {
		return true
	}

	for _, f := range filters {
		if !f(eh) {
			return false
		}
	}

	return true
}
