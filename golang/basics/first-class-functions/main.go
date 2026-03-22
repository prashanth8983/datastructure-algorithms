package main

import (
	"fmt"
	"math"
	"strings"
)

// ============ FIRST-CLASS FUNCTIONS ============
// In Go, functions are first-class: assign to variables, pass as arguments, return from functions

// ============ FUNCTION TYPES ============

// TODO: Declare a function type "transformer" that takes a string and returns a string
type transformer func(string) string

// TODO: Declare a function type "mathOp" that takes two float64s and returns float64
type mathOp func(float64, float64) float64

// ============ FUNCTIONS TO USE ============

func double(x float64) float64 {
	return x * 2
}

func square(x float64) float64 {
	return x * x
}

func add(a, b float64) float64  { return a + b }
func mult(a, b float64) float64 { return a * b }

// ============ PASSING FUNCTIONS AS PARAMETERS ============

// TODO: Write apply() that takes a []float64 and a func(float64) float64,
//       returns a new slice with the function applied to each element
func apply(nums []float64, fn func(float64) float64) []float64 {
	result := make([]float64, len(nums))
	for i, v := range nums {
		result[i] = fn(v)
	}
	return result
}

// TODO: Write applyOp() that takes two float64s and a mathOp, returns the result
func applyOp(a, b float64, op mathOp) float64 {
	return op(a, b)
}

// TODO: Write transform() that takes a []string and a transformer, returns transformed slice
func transform(words []string, fn transformer) []string {
	result := make([]string, len(words))
	for i, w := range words {
		result[i] = fn(w)
	}
	return result
}

// ============ RETURNING FUNCTIONS (CLOSURES) ============

// TODO: Write makeMultiplier() that takes an int and returns a func(int) int
//       The returned function multiplies its argument by the captured value
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// TODO: Write makeGreeter() that takes a greeting string and returns a func(string) string
//       The returned function prepends the greeting to the name
func makeGreeter(greeting string) func(string) string {
	return func(name string) string {
		return greeting + ", " + name + "!"
	}
}

// TODO: Write makeCounter() that returns a func() int
//       Each call increments and returns the count (closure captures the counter)
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// ============ CALIBRATE PATTERN (from the book) ============
// A function that wraps another function with an adjustment

// TODO: Write calibrate() that takes a sensor func() float64 and an offset float64
//       Returns a new func() float64 that adds the offset to the sensor reading
func calibrate(sensor func() float64, offset float64) func() float64 {
	return func() float64 {
		return sensor() + offset
	}
}

func main() {
	// ============ ASSIGNING FUNCTIONS TO VARIABLES ============

	// TODO: Assign the double function to a variable called fn (no parentheses!)
	fn := double
	fmt.Println("double(5):", fn(5))

	// TODO: Reassign fn to the square function
	fn = square
	fmt.Println("square(5):", fn(5))

	// ============ PASSING FUNCTIONS AS ARGUMENTS ============

	nums := []float64{1, 2, 3, 4, 5}

	// TODO: Use apply() with double
	doubled := apply(nums, double)
	fmt.Println("\nDoubled:", doubled)

	// TODO: Use apply() with square
	squared := apply(nums, square)
	fmt.Println("Squared:", squared)

	// TODO: Use apply() with math.Sqrt (standard library function!)
	sqrts := apply(nums, math.Sqrt)
	fmt.Println("Sqrt:", sqrts)

	// TODO: Use applyOp() with add and mult
	fmt.Println("\n3 + 4 =", applyOp(3, 4, add))
	fmt.Println("3 * 4 =", applyOp(3, 4, mult))

	// ============ USING FUNCTION TYPES ============

	// TODO: Use transform() with strings.ToUpper
	words := []string{"hello", "gopher", "world"}
	upper := transform(words, strings.ToUpper)
	fmt.Println("\nUpper:", upper)

	// TODO: Use transform() with strings.ToLower
	lower := transform([]string{"GO", "IS", "FUN"}, strings.ToLower)
	fmt.Println("Lower:", lower)

	// ============ ANONYMOUS FUNCTIONS ============

	// TODO: Pass an anonymous function to apply() that triples each value
	tripled := apply(nums, func(x float64) float64 {
		return x * 3
	})
	fmt.Println("\nTripled:", tripled)

	// TODO: Declare and immediately invoke an anonymous function that prints a message
	func() {
		fmt.Println("Anonymous function executed!")
	}()

	// TODO: Anonymous function with a parameter
	func(msg string) {
		fmt.Println("Message:", msg)
	}("Hello from anonymous")

	// ============ CLOSURES - RETURNING FUNCTIONS ============

	// TODO: Use makeMultiplier to create a tripler and a doubler
	tripler := makeMultiplier(3)
	doubler := makeMultiplier(2)
	fmt.Println("\nTripler(10):", tripler(10))
	fmt.Println("Doubler(10):", doubler(10))

	// TODO: Use makeGreeter to create different greeters
	helloGreet := makeGreeter("Hello")
	hiGreet := makeGreeter("Hi")
	fmt.Println("\n" + helloGreet("Gopher"))
	fmt.Println(hiGreet("World"))

	// TODO: Use makeCounter - call it multiple times to see the state persist
	counter := makeCounter()
	fmt.Println("\nCounter:", counter()) // 1
	fmt.Println("Counter:", counter())   // 2
	fmt.Println("Counter:", counter())   // 3

	// TODO: Create a second counter to show they are independent
	counter2 := makeCounter()
	fmt.Println("Counter2:", counter2()) // 1 (independent!)

	// ============ CALIBRATE PATTERN ============

	// TODO: Create a fake sensor and calibrate it with an offset of 5
	fakeSensor := func() float64 { return 100.0 }
	calibrated := calibrate(fakeSensor, 5)
	fmt.Println("\nSensor:", fakeSensor())
	fmt.Println("Calibrated:", calibrated())

	_ = fmt.Println // remove once used
}
