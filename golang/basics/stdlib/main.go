package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// ============ MATH ============
	fmt.Println("=== Math ===")

	// TODO: Print math.MaxInt and math.MinInt

	// TODO: Get absolute value of -5 (note: math.Abs takes float64, no int version)
	//       Use: int(math.Abs(float64(-5)))

	// TODO: Use built-in min() and max() (Go 1.21+) with integers
	//       Print min(3, 7) and max(3, 7)

	// ============ STRING FORMATTING ============
	fmt.Println("\n=== String Formatting ===")

	x, y := 10, 20

	// TODO: Create a formatted string using fmt.Sprintf("%d-%d", x, y) and print it

	// TODO: Create a formatted string with padding: fmt.Sprintf("%05d", 42) and print

	_ = x
	_ = y

	// ============ STRCONV ============
	fmt.Println("\n=== Strconv ===")

	// TODO: Convert 42 to string using strconv.Itoa and print

	// TODO: Convert "123" to int using strconv.Atoi (handle the error or use _)

	// TODO: Convert 3.14 to string using fmt.Sprintf("%.2f", 3.14) and print

	// ============ STRINGS PACKAGE ============
	fmt.Println("\n=== Strings Package ===")

	// TODO: Convert "Hello World" to uppercase using strings.ToUpper

	// TODO: Convert "Hello World" to lowercase using strings.ToLower

	// TODO: Trim spaces from "  hello  " using strings.TrimSpace

	// TODO: Replace "hello world" -> "hello Go" using strings.Replace

	// TODO: Count occurrences of "l" in "hello world" using strings.Count

	// TODO: Find index of "world" in "hello world" using strings.Index

	_ = math.MaxInt      // remove once used
	_ = strconv.Itoa     // remove once used
	_ = strings.ToUpper  // remove once used
	_ = fmt.Println      // remove once used
}
