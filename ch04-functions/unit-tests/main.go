package main

func getMonthlyPrice(tier string) int {
	price := 0.0
	switch tier {
	case "basic":
		price = 100.0 * 100.0
	case "premium":
		price = 150.0 * 100.0
	case "enterprise":
		price = 500.0 * 100.0
	default:
		price = 0.0
	}
	
	return int(price)
}

