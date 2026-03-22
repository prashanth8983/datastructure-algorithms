package main

import (
	"fmt"
	"math"
)

// ============ DEFINING INTERFACES ============

// TODO: Define an interface Shape with two methods:
//       Area() float64
//       Perimeter() float64

// TODO: Define a struct Circle with field Radius float64

// TODO: Define a struct Rectangle with fields Width, Height float64

// TODO: Implement Area() and Perimeter() on Circle
//       Area = π * r²,  Perimeter = 2 * π * r

// TODO: Implement Area() and Perimeter() on Rectangle
//       Area = w * h,  Perimeter = 2 * (w + h)

// ============ INTERFACE AS PARAMETER ============

// TODO: Write a function printShapeInfo(s Shape) that prints area and perimeter
//       func printShapeInfo(s Shape) {
//           fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
//       }

// ============ EMPTY INTERFACE ============

// TODO: Write a function printAnything(val interface{}) that prints any value
//       (interface{} is like Object in Java, or any in TypeScript)
//       In Go 1.18+, you can also use "any" as an alias

// ============ TYPE ASSERTION ============

// TODO: Write a function describe(i interface{}) that uses type assertion:
//       func describe(i interface{}) {
//           switch v := i.(type) {
//           case int:
//               fmt.Println("int:", v)
//           case string:
//               fmt.Println("string:", v)
//           case Shape:
//               fmt.Println("shape with area:", v.Area())
//           default:
//               fmt.Println("unknown type")
//           }
//       }

// ============ STRINGER INTERFACE ============
// Implementing fmt.Stringer (like toString() in Java)

// TODO: Define a struct Point with X, Y int
// TODO: Implement String() string on Point
//       func (p Point) String() string {
//           return fmt.Sprintf("(%d, %d)", p.X, p.Y)
//       }

// ============ ERROR INTERFACE ============
// error is just an interface: Error() string

// TODO: Define a custom error type DivisionError with field Message string
// TODO: Implement Error() string on DivisionError
// TODO: Write a function safeDivide(a, b float64) (float64, error)
//       that returns DivisionError if b == 0

func main() {
	// TODO: Create a Circle{Radius: 5} and Rectangle{Width: 3, Height: 4}
	// TODO: Call printShapeInfo on both

	// TODO: Call printAnything with 42, "hello", and true

	// TODO: Call describe with 42, "hello", and your Circle

	// TODO: Create a Point{3, 4} and print it (String() is called automatically)

	// TODO: Call safeDivide(10, 0) and handle the error
	//       result, err := safeDivide(10, 0)
	//       if err != nil { fmt.Println("Error:", err) }

	_ = math.Pi     // remove once used
	_ = fmt.Println // remove once used
}
