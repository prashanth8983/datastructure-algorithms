package main

import "fmt"

func main() {
	// ============ SLICE BASICS ============

	// TODO: Create a slice s := []int{1, 2, 3}

	// TODO: Append 4 to s (push)

	// TODO: Append 5, 6, 7 to s (push multiple)

	// TODO: Create another slice other := []int{8, 9}
	// TODO: Append other to s using spread (...)

	fmt.Println("After appends:", s)

	// TODO: Pop last element from s (s = s[:len(s)-1])

	// TODO: Dequeue/shift first element from s (s = s[1:])

	fmt.Println("After pop and shift:", s)

	// ============ MAKE ============

	// TODO: Create a slice of length 5 using make (all zeros)

	// TODO: Create a slice of length 0, capacity 5 using make

	fmt.Println("Make length 5:", zeros)
	fmt.Println("Make cap 5:", "len=", len(withCap), "cap=", cap(withCap))

	// ============ 2D SLICE ============
	rows, cols := 3, 4

	// TODO: Create a 2D slice (grid) with 'rows' rows and 'cols' columns
	//       Use make for outer slice, then loop to make each inner slice

	grid[1][2] = 42
	fmt.Println("2D grid:", grid)

	// ============ SLICE GOTCHA: SHARED MEMORY ============
	a := []int{1, 2, 3, 4}
	b := a[1:3] // [2, 3] - shares backing array with a!
	b[0] = 99
	fmt.Println("a after modifying b:", a) // a is affected!

	// TODO: Create an independent copy of a using make + copy

	// TODO: OR create an independent copy using append([]int{}, a...)

	c[0] = 777
	fmt.Println("a after modifying independent copy:", a) // a is NOT affected

	_ = fmt.Println // remove once used
}
