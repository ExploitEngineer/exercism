package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinut calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	cardsPerHour := float64(productionRate) * (successRate / 100)

	carsPerMinute := cardsPerHour / 60

	return int(carsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsOfTen := carsCount / 10

	individualCars := carsCount % 10

	finalAnswer := (groupsOfTen * 95000) + (individualCars * 10000)

	return uint(finalAnswer)
}
