package main

import (
	"fmt"
	"sort"
)

func main() {
	// ============ BASIC SORTING ============

	// TODO: Sort []int{5, 2, 8, 1, 9} in ascending order using sort.Ints

	fmt.Println("Ascending:", arr)

	// TODO: Sort []string{"banana", "apple", "cherry"} using sort.Strings

	fmt.Println("Strings:", strs)

	// TODO: Sort arr in descending order using sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	fmt.Println("Descending:", arr)

	// ============ CUSTOM SORT ============

	// Sort intervals by start time, then by end time
	intervals := [][]int{{1, 3}, {2, 4}, {1, 2}, {3, 5}, {2, 3}}

	// TODO: Sort intervals using sort.Slice
	//       Primary sort by [0] ascending, secondary sort by [1] ascending

	fmt.Println("Sorted intervals:", intervals)

	// ============ SORT STRUCTS ============

	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Carol", 35},
		{"Dave", 25},
	}

	// TODO: Sort people by age ascending, then by name alphabetically for same age

	fmt.Println("Sorted people:", people)

	_ = sort.Ints // remove once used
}
