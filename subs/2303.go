package subs

func calculateTax(brackets [][]int, income int) float64 {
	var r float64
	var lastTaxBound int
	for _, bracket := range brackets {

		taxedIncome := bracket[0] - lastTaxBound
		if income <= taxedIncome {
			taxedIncome = income
		}

		r += float64(taxedIncome) * float64(bracket[1]) / 100.0

		lastTaxBound = bracket[0]
		income -= taxedIncome
		if income <= 0 {
			break
		}

	}
	return r
}
