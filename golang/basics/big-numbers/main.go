package main

import (
	"fmt"
	"math/big"
)

func main() {
	// ============ THE CEILING: WHEN int64 ISN'T ENOUGH ============
	// int64 max: 9,223,372,036,854,775,807 (~9.2 quintillion)
	// What about distances to Andromeda? Factorial of 100? Crypto keys?

	// TODO: Demonstrate the limit of int64 with light speed calculation
	const lightSpeed = 299792 // km/s
	const secondsPerDay = 86400
	var distance int64 = 41.3e12 // Alpha Centauri in km
	days := distance / int64(lightSpeed) / int64(secondsPerDay)
	fmt.Println("Days to Alpha Centauri:", days)

	// ============ big.Int - ARBITRARY PRECISION INTEGERS ============

	// TODO: Create a big.Int using NewInt (for small values)
	a := big.NewInt(100)
	fmt.Println("\nbig.Int from NewInt:", a)

	// TODO: Create a big.Int from a string using SetString (for large values)
	//       Use the distance to Andromeda: 24,000,000,000,000,000,000 km (24 quintillion)
	andromeda, _ := new(big.Int).SetString("24000000000000000000", 10)
	fmt.Println("Andromeda distance:", andromeda, "km")

	// ============ big.Int ARITHMETIC ============
	// big.Int uses methods, not operators: Add, Sub, Mul, Div, Mod

	// TODO: Calculate 100! (factorial of 100) using big.Int
	factorial := big.NewInt(1)
	for i := 2; i <= 100; i++ {
		factorial.Mul(factorial, big.NewInt(int64(i)))
	}
	fmt.Println("\n100! =", factorial)

	// TODO: Add two big.Ints
	x := big.NewInt(1000000000000)
	y := big.NewInt(2000000000000)
	sum := new(big.Int).Add(x, y)
	fmt.Println("\nSum:", sum)

	// TODO: Multiply two big.Ints
	product := new(big.Int).Mul(x, y)
	fmt.Println("Product:", product)

	// TODO: Divide two big.Ints (integer division)
	quotient := new(big.Int).Div(product, x)
	fmt.Println("Quotient:", quotient)

	// TODO: Modulo with big.Int
	mod := new(big.Int).Mod(big.NewInt(17), big.NewInt(5))
	fmt.Println("17 mod 5:", mod)

	// ============ big.Int COMPARISON ============
	// Cmp returns -1, 0, or 1

	// TODO: Compare two big.Ints
	cmp := x.Cmp(y)
	fmt.Println("\nx.Cmp(y):", cmp) // -1 (x < y)

	// ============ big.Int EXPONENTIATION ============

	// TODO: Calculate 2^256 (common in crypto)
	two := big.NewInt(2)
	exp := new(big.Int).Exp(two, big.NewInt(256), nil) // nil = no modulus
	fmt.Println("\n2^256 =", exp)

	// TODO: Calculate 2^256 with a modulus (modular exponentiation, used in RSA)
	modulus := big.NewInt(1000000007)
	modExp := new(big.Int).Exp(two, big.NewInt(256), modulus)
	fmt.Println("2^256 mod 1e9+7 =", modExp)

	// ============ big.Float - ARBITRARY PRECISION FLOATS ============

	// TODO: Create a big.Float with high precision for pi
	pi := new(big.Float).SetPrec(200).SetFloat64(3.14159265358979323846)
	fmt.Println("\nbig.Float pi:", pi)

	// TODO: Arithmetic with big.Float
	bf1 := new(big.Float).SetFloat64(1.0)
	bf3 := new(big.Float).SetFloat64(3.0)
	third := new(big.Float).SetPrec(100).Quo(bf1, bf3)
	fmt.Println("1/3 =", third)

	// ============ CONSTANTS AND big.Int ============
	// Go constants have unlimited precision (evaluated at compile time)

	// Go constants have unlimited precision at compile time
	// They only need to fit in a type when assigned to a variable
	// Example: this constant exceeds int64 but is valid in constant expressions
	const bigConst = 24000000000000000000
	const half = bigConst / 2 // constant arithmetic works at arbitrary precision
	// fmt.Println(half) would fail - overflows int64 at runtime
	// But you can use it with big.Int:
	halfBig, _ := new(big.Int).SetString("12000000000000000000", 10)
	fmt.Println("\nConstant half (via big.Int):", halfBig)

	// ============ PRACTICAL: FIBONACCI WITH big.Int ============

	// TODO: Calculate the 200th Fibonacci number using big.Int
	fib := bigFib(200)
	fmt.Println("\nFib(200) =", fib)

	_ = fmt.Println // remove once used
}

// TODO: Write bigFib(n) that returns the nth Fibonacci number as *big.Int
func bigFib(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}
	a := big.NewInt(0)
	b := big.NewInt(1)
	for i := 2; i <= n; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return b
}
