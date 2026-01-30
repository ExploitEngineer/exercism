// Package weather provides tools to get
// information about
// current city weather condition.
package weather

var (
	// CurrentCondition variable will store the weather condition according to the weather.
	CurrentCondition string
	// CurrentLocation variable will take your current location to get the weather information.
	CurrentLocation string
)

// Forecast returns a string which tell the weather condition of the the specified city. It will take two parameters city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
