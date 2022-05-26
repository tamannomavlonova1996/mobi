package tom

func Calculate(amount float64) float64 {
	comission := 0.005

	if amount < 5000 {
		return 0
	}
	result := 0.0
	result = amount * comission
	return result
}
