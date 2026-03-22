package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ============ VARIABLES ============
	// TODO: Declare x using short declaration and assign 10
	x := 10
	// TODO: Declare y with explicit type int and assign 20
	var y int = 20
	// TODO: Declare z as int (zero value)
	var z int
	// TODO: Declare a, b using multiple short declaration with values 1, 2
	a ,b := 1,2
	fmt.Println("Variables:", x, y, z, a, b)

	// ============ CONSTANTS ============
	// TODO: Declare a constant PI = 3.14
	const PI = 3.14
	// TODO: Declare a const block with MOD = 1e9 + 7 and INF = int(1e18)
	const (
		MOD = 1e9+7
		INF = int(1e18)
	)
	fmt.Println("Constants:", PI, MOD, INF)

	// ============ TYPE CONVERSIONS ============
	// Go has NO implicit casting - you must convert explicitly

	// TODO: Convert x (int) to float64 and store in f
	f := float64(x)
	// TODO: Convert f back to int and store in n
	n := int(f)
	// TODO: Convert rune 65 to string and store in s (should be "A")
	s := string(rune(65))
	// TODO: Convert int 42 to string using strconv.Itoa, store in num
	num := strconv.Itoa(42)
	// TODO: Convert string "42" to int using strconv.Atoi, store in val (ignore error with _)

	val, _ := strconv.Atoi("42")
	_ = strconv.Itoa // remove this line once you use strconv

	fmt.Println("Conversions:", f, n, s, num, val)
}
