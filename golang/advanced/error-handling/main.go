package main

import (
	"errors"
	"fmt"
	"strconv"
)

// ============ BASIC ERROR HANDLING ============

// TODO: Write a function divide(a, b float64) (float64, error)
//       Return fmt.Errorf("cannot divide by zero") if b == 0
//       Otherwise return a/b, nil

// ============ CUSTOM ERROR TYPE ============

// TODO: Define a struct ValidationError with fields:
//       Field   string
//       Message string

// TODO: Implement Error() string on ValidationError
//       Return: fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)

// TODO: Write a function validateAge(age int) error
//       Return &ValidationError{"age", "must be positive"} if age < 0
//       Return &ValidationError{"age", "unrealistic age"} if age > 150
//       Return nil otherwise

// ============ WRAPPING ERRORS ============

// TODO: Write a function parseAndDouble(s string) (int, error)
//       Use strconv.Atoi to parse s
//       If err != nil, wrap it: fmt.Errorf("parseAndDouble failed: %w", err)
//       Otherwise return parsed * 2, nil

// ============ ERRORS.IS AND ERRORS.AS ============

// TODO: Define a sentinel error:
//       var ErrNotFound = errors.New("not found")

// TODO: Write a function findUser(id int) (string, error)
//       Return "", fmt.Errorf("findUser: %w", ErrNotFound) if id != 1
//       Return "Alice", nil if id == 1

func main() {
	// ============ BASIC ============
	fmt.Println("=== Basic Error Handling ===")

	// TODO: Call divide(10, 0) and handle the error
	//       if err != nil { fmt.Println("Error:", err) }
	// TODO: Call divide(10, 3) and print the result

	// ============ CUSTOM ERROR ============
	fmt.Println("\n=== Custom Error ===")

	// TODO: Call validateAge(-5)
	// TODO: Use errors.As to check if it's a *ValidationError and print the Field
	//       var ve *ValidationError
	//       if errors.As(err, &ve) {
	//           fmt.Println("Invalid field:", ve.Field)
	//       }

	// ============ WRAPPING ============
	fmt.Println("\n=== Wrapped Errors ===")

	// TODO: Call parseAndDouble("abc") and print the error
	// TODO: Call parseAndDouble("21") and print the result

	// ============ ERRORS.IS ============
	fmt.Println("\n=== errors.Is ===")

	// TODO: Call findUser(99)
	// TODO: Use errors.Is(err, ErrNotFound) to check the error type

	_ = errors.New    // remove once used
	_ = strconv.Atoi  // remove once used
	_ = fmt.Println   // remove once used
}
