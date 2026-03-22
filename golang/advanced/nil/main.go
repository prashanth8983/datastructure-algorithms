package main

import (
	"fmt"
	"strings"
)

// ============ NIL IN GO ============
// nil is the zero value for pointers, slices, maps, channels, functions, and interfaces
// Each type of nil behaves differently - understanding nil prevents panics

func main() {
	// ============ NIL POINTERS ============

	// TODO: Declare a nil pointer to int
	var p *int
	fmt.Println("Nil pointer:", p)       // <nil>
	fmt.Println("Is nil?:", p == nil)     // true

	// TODO: Dereferencing a nil pointer would panic!
	//       Always check before dereferencing
	if p != nil {
		fmt.Println("Value:", *p)
	} else {
		fmt.Println("Cannot dereference nil pointer")
	}

	// TODO: Assign a value and check again
	val := 42
	p = &val
	fmt.Println("After assign:", *p)

	// ============ NIL SLICES ============
	// A nil slice has length 0, capacity 0, and no underlying array
	// BUT you can still append to it! (Go handles this gracefully)

	// TODO: Declare a nil slice
	var s []int
	fmt.Println("\nNil slice:", s)
	fmt.Println("Length:", len(s))     // 0
	fmt.Println("Is nil?:", s == nil)  // true

	// TODO: Append to nil slice (this works!)
	s = append(s, 1, 2, 3)
	fmt.Println("After append:", s)
	fmt.Println("Is nil now?:", s == nil) // false

	// TODO: Range over nil slice (safe - just iterates zero times)
	var empty []int
	fmt.Print("Range nil slice: ")
	for _, v := range empty {
		fmt.Print(v)
	}
	fmt.Println("(nothing printed)")

	// ============ NIL MAPS ============
	// A nil map can be READ from (returns zero values) but CANNOT be written to (panic!)

	// TODO: Declare a nil map
	var m map[string]int
	fmt.Println("\nNil map:", m)
	fmt.Println("Is nil?:", m == nil) // true

	// TODO: Reading from nil map is safe (returns zero value)
	fmt.Println("Read nil map:", m["key"]) // 0

	// TODO: Two-value form also safe
	v, ok := m["key"]
	fmt.Println("Value:", v, "Exists:", ok) // 0, false

	// TODO: Writing to nil map would PANIC!
	//       m["key"] = 1  // panic: assignment to entry in nil map
	//       Always initialize with make() or literal before writing
	m = make(map[string]int)
	m["key"] = 1
	fmt.Println("After init:", m)

	// ============ NIL FUNCTIONS ============

	// TODO: Declare a nil function variable
	var fn func(int) int
	fmt.Println("\nNil func:", fn == nil) // true

	// TODO: Calling a nil function would panic!
	//       fn(5) // panic: call of nil function
	//       Always check before calling
	if fn != nil {
		fmt.Println("Result:", fn(5))
	} else {
		fmt.Println("Cannot call nil function")
	}

	// TODO: Assign a function and use it
	fn = func(x int) int { return x * 2 }
	fmt.Println("After assign:", fn(5)) // 10

	// ============ NIL INTERFACES ============
	// An interface is nil only when BOTH its type and value are nil
	// This is a common source of bugs!

	// TODO: Demonstrate nil interface
	var err error
	fmt.Println("\nNil interface:", err)
	fmt.Println("Is nil?:", err == nil) // true

	// TODO: Demonstrate the nil interface trap
	//       A non-nil interface can hold a nil pointer - and it won't equal nil!
	var strPtr *myError
	err = strPtr // non-nil interface wrapping a nil *myError
	fmt.Println("Nil pointer in interface:", err)
	fmt.Println("Is nil?:", err == nil)             // FALSE! (the interface itself is not nil)
	fmt.Println("Type is not nil, value is nil")

	// ============ SAFE NIL PATTERNS ============

	// TODO: Write a safe string joiner that handles nil slices
	fmt.Println("\nSafe join nil:", safeJoin(nil, ","))
	fmt.Println("Safe join empty:", safeJoin([]string{}, ","))
	fmt.Println("Safe join values:", safeJoin([]string{"a", "b", "c"}, ","))

	// TODO: Write a function that safely gets from a potentially nil map
	data := map[string]int{"x": 10}
	fmt.Println("\nSafe get 'x':", safeGet(data, "x", -1))
	fmt.Println("Safe get 'y':", safeGet(data, "y", -1))
	fmt.Println("Safe get nil:", safeGet(nil, "x", -1))

	// ============ NIL CHANNELS ============
	// Sending or receiving on a nil channel blocks forever
	// Useful for disabling select cases

	// TODO: Demonstrate nil channel
	var ch chan int
	fmt.Println("\nNil channel:", ch == nil) // true
	// ch <- 1   // blocks forever
	// <-ch      // blocks forever

	// In a select, a nil channel case is never selected (used to disable branches)
	fmt.Println("Nil channels are used to disable select cases")

	_ = fmt.Println // remove once used
}

// Custom error type for demonstrating nil interface trap
type myError struct {
	msg string
}

func (e *myError) Error() string {
	if e == nil {
		return "nil error"
	}
	return e.msg
}

// TODO: Implement safeJoin - joins strings with separator, handles nil/empty slice
func safeJoin(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	return strings.Join(strs, sep)
}

// TODO: Implement safeGet - gets value from map with default, handles nil map
func safeGet(m map[string]int, key string, defaultVal int) int {
	if m == nil {
		return defaultVal
	}
	if v, ok := m[key]; ok {
		return v
	}
	return defaultVal
}
