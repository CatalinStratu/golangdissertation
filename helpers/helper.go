package helpers

func CalculateValueFromPercent(total float64, percent float64) float64 {
	if percent < 0 || percent > 100 {
		panic("Percentage must be between 0 and 100")
	}

	return (percent / 100) * total
}
