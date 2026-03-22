package main

import "fmt"

// ============ POINTER BASICS ============

func main() {
	fmt.Println("=== Pointer Basics ===")

	// TODO: Declare x := 42
	// TODO: Create a pointer to x: p := &x
	// TODO: Print the pointer (address) and the dereferenced value (*p)
	// TODO: Modify x through the pointer: *p = 100
	// TODO: Print x to confirm it changed

	// ============ POINTERS IN FUNCTIONS ============
	fmt.Println("\n=== Pointers in Functions ===")

	// TODO: Write a function increment(val *int) that does *val++
	//       (modifies the original value, like pass by reference)

	// TODO: Call increment(&x) and print x

	// TODO: Write a function tryIncrement(val int) that does val++
	//       (does NOT modify original — pass by value)

	// TODO: Call tryIncrement(x) and print x to show it didn't change

	// ============ POINTER VS VALUE RECEIVERS ============
	fmt.Println("\n=== Pointer vs Value Receivers ===")

	// TODO: Define a struct Counter with field N int

	// TODO: Implement a method with pointer receiver:
	//       func (c *Counter) Increment() { c.N++ }
	//       (this modifies the original struct)

	// TODO: Implement a method with value receiver:
	//       func (c Counter) Value() int { return c.N }
	//       (this works on a copy)

	// TODO: Create a Counter, call Increment 3 times, print Value()

	// ============ NIL POINTERS ============
	fmt.Println("\n=== Nil Pointers ===")

	// TODO: Declare a nil pointer: var p *int
	// TODO: Check if p is nil and print result
	// TODO: Show safe pattern:
	//       if p != nil { fmt.Println(*p) } else { fmt.Println("p is nil") }

	// ============ NEW KEYWORD ============
	fmt.Println("\n=== New ===")

	// TODO: Allocate with new: p := new(int)  // returns *int, zero-valued
	// TODO: Print *p (should be 0)
	// TODO: Set *p = 42 and print again

	_ = fmt.Println // remove once used
}
