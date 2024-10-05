package calc

type Discount interface {
	Amount(amount, discount float64) float64
}

type MinPriceDiscount struct {
	calculator Calculator
	minAmount  float64
}

func NewMinPriceDiscount(minAmount float64, calculator Calculator) Discount {
	return MinPriceDiscount{
		calculator: calculator,
		minAmount:  minAmount,
	}
}

func (d MinPriceDiscount) Amount(amount, discount float64) float64 {
	if amount < d.minAmount {
		return amount
	}

	if amount < discount {
		return amount
	}
	return float64(d.calculator.Subtract(int(amount), int(discount)))
}
