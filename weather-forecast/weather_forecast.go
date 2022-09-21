// Package weather: contains tools for weather forecasting relative to cities.
package weather

// CurrentCondition: contains the currently set condition. Its value is set by the last condition that has been passed to the Forecast function.
var CurrentCondition string

// CurrentLocation: contains the currently set condition. Its value is set by the last location that has been passed to the Forecast function.
var CurrentLocation string

// Forecast: returns a string containing information on the current weather condition of a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
