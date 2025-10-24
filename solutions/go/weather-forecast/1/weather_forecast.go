// Package weather provides tools to
// print weather forecast for a specific location.
package weather


var (
	// CurrentCondition represents the current weather condition.
	CurrentCondition string
	// CurrentLocation represents the location for which forecast is provided.
	CurrentLocation  string
)

// Forecast returns a string representing weather forecast at specific location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
