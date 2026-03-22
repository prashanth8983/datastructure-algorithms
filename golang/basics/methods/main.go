package main

import "fmt"

// ============ CUSTOM TYPES ============
// Go lets you declare new types based on existing ones
// This adds meaning and enables methods

// TODO: Declare a custom type "celsius" based on float64
type celsius float64

// TODO: Declare a custom type "fahrenheit" based on float64
type fahrenheit float64

// TODO: Declare a custom type "kelvin" based on float64
type kelvin float64

// ============ METHODS ON CUSTOM TYPES ============
// Methods are functions with a receiver (the type they belong to)
// Syntax: func (receiver) methodName() returnType

// TODO: Write a method on kelvin that converts to celsius
//       Formula: celsius = kelvin - 273.15
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

// TODO: Write a method on kelvin that converts to fahrenheit
//       Formula: fahrenheit = (kelvin - 273.15) * 9.0/5.0 + 32.0
func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit((k-273.15)*9.0/5.0 + 32.0)
}

// TODO: Write a method on celsius that converts to kelvin
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

// TODO: Write a method on celsius that converts to fahrenheit
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*9.0/5.0 + 32.0)
}

// TODO: Write a method on fahrenheit that converts to celsius
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

// TODO: Write a method on fahrenheit that converts to kelvin
func (f fahrenheit) kelvin() kelvin {
	return kelvin((f-32.0)*5.0/9.0 + 273.15)
}

// ============ CUSTOM TYPE WITH MULTIPLE METHODS ============

// TODO: Declare a custom type "meter" based on float64
type meter float64

// TODO: Write a method to convert meter to kilometers
func (m meter) km() float64 {
	return float64(m) / 1000.0
}

// TODO: Write a method to convert meter to miles (1 mile = 1609.34 meters)
func (m meter) miles() float64 {
	return float64(m) / 1609.34
}

// ============ METHODS WITH PARAMETERS ============

// TODO: Declare a custom type "coordinate" as a struct with lat, long float64
type coordinate struct {
	lat  float64
	long float64
}

// TODO: Write a method description() on coordinate that returns a formatted string
func (c coordinate) description() string {
	return fmt.Sprintf("(%.2f, %.2f)", c.lat, c.long)
}

// TODO: Write a method isNorthernHemisphere() that returns true if lat > 0
func (c coordinate) isNorthernHemisphere() bool {
	return c.lat > 0
}

func main() {
	// ============ USING CUSTOM TYPE METHODS ============

	// TODO: Create a kelvin value of 294.0 and convert to celsius and fahrenheit
	var k kelvin = 294.0
	fmt.Printf("%.1f K = %.1f C = %.1f F\n", k, k.celsius(), k.fahrenheit())

	// TODO: Create a celsius value of 21.0 and convert to kelvin and fahrenheit
	var c celsius = 21.0
	fmt.Printf("%.1f C = %.1f K = %.1f F\n", c, c.kelvin(), c.fahrenheit())

	// TODO: Create a fahrenheit value of 69.8 and convert to celsius and kelvin
	var f fahrenheit = 69.8
	fmt.Printf("%.1f F = %.1f C = %.1f K\n", f, f.celsius(), f.kelvin())

	// TODO: Chain conversions: kelvin -> celsius -> fahrenheit
	chained := k.celsius().fahrenheit()
	fmt.Printf("Chained: %.1f K -> %.1f F\n", k, chained)

	// ============ USING METER METHODS ============

	// TODO: Create a distance of 42195 meters (marathon) and print in km and miles
	marathon := meter(42195)
	fmt.Printf("\nMarathon: %.0f m = %.2f km = %.2f miles\n", marathon, marathon.km(), marathon.miles())

	// ============ USING STRUCT METHODS ============

	// TODO: Create a coordinate for New York (40.7128, -74.0060)
	nyc := coordinate{lat: 40.7128, long: -74.0060}
	fmt.Printf("\nNYC: %s, Northern: %v\n", nyc.description(), nyc.isNorthernHemisphere())

	// TODO: Create a coordinate for Sydney (-33.8688, 151.2093)
	sydney := coordinate{lat: -33.8688, long: 151.2093}
	fmt.Printf("Sydney: %s, Northern: %v\n", sydney.description(), sydney.isNorthernHemisphere())

	// ============ KEY DIFFERENCE: FUNCTION VS METHOD ============
	// Function: kelvinToCelsius(k)  - standalone function
	// Method:   k.celsius()         - called on the value with dot notation
	// Methods associate behavior with types, making code more readable

	_ = fmt.Println // remove once used
}
