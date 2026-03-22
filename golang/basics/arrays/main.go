package main

import "fmt"

func main() {
	// ============ DECLARING ARRAYS (fixed-size, unlike slices) ============
	// In Go, arrays have a FIXED length that is part of the type
	// [8]string and [5]string are DIFFERENT types

	// TODO: Declare an array of 8 strings called planets (zero-valued)
	var planets [8]string

	// TODO: Assign "Mercury", "Venus", "Earth" to indices 0, 1, 2
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	// TODO: Print the length of planets (should be 8, not 3!)
	fmt.Println("Length:", len(planets))

	// TODO: Print planets[3] to see the zero value for strings (empty string)
	fmt.Println("Zero value:", planets[3] == "")

	// ============ COMPOSITE LITERALS ============
	// Initialize arrays inline with values

	// TODO: Declare dwarfs as [5]string with "Ceres", "Pluto", "Haumea", "Makemake", "Eris"
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Println("Dwarfs:", dwarfs)

	// TODO: Declare allPlanets using [...] (compiler counts elements)
	//       with all 8 planets: Mercury, Venus, Earth, Mars, Jupiter, Saturn, Uranus, Neptune
	allPlanets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println("All planets:", allPlanets)
	fmt.Println("Count:", len(allPlanets))

	// ============ ITERATING ARRAYS ============

	// TODO: Loop through dwarfs using classic for loop, print index and value
	fmt.Println("\n--- Classic loop ---")
	for i := 0; i < len(dwarfs); i++ {
		fmt.Println(i, dwarfs[i])
	}

	// TODO: Loop through dwarfs using range, print index and value
	fmt.Println("\n--- Range loop ---")
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}

	// ============ ARRAYS ARE VALUES (copied on assignment) ============
	// Unlike slices, arrays are VALUE types - assigning copies the entire array

	// TODO: Copy allPlanets into a new variable called backup
	backup := allPlanets

	// TODO: Modify backup[0] to "MODIFIED"
	backup[0] = "MODIFIED"

	// TODO: Print both to prove original is unchanged
	fmt.Println("\nOriginal[0]:", allPlanets[0]) // still "Mercury"
	fmt.Println("Backup[0]:", backup[0])         // "MODIFIED"

	// ============ ARRAYS VS SLICES ============
	// Key differences:
	// - Arrays: fixed size, value type (copied), size is part of the type
	// - Slices: dynamic size, reference type (shared backing array)

	// TODO: Declare an array of 3 ints: [3]int{10, 20, 30}
	arr := [3]int{10, 20, 30}

	// TODO: Create a slice FROM the array using arr[:] (this creates a view into the array)
	slc := arr[:]

	// TODO: Modify slc[0] = 99, then print arr[0] to see it changed (shared memory!)
	slc[0] = 99
	fmt.Println("\nArray after slice modify:", arr[0]) // 99 - array changed!

	// TODO: Create a slice from index 1 to 3 of allPlanets
	innerPlanets := allPlanets[1:3]
	fmt.Println("Inner planets:", innerPlanets) // [Venus Earth]

	// ============ 2D ARRAYS ============

	// TODO: Declare a 3x3 grid as [3][3]int
	var grid [3][3]int

	// TODO: Set grid[0][0] = 1, grid[1][1] = 5, grid[2][2] = 9 (diagonal)
	grid[0][0] = 1
	grid[1][1] = 5
	grid[2][2] = 9

	// TODO: Print the grid row by row
	fmt.Println("\n--- 3x3 Grid ---")
	for _, row := range grid {
		fmt.Println(row)
	}

	// ============ COMPARING ARRAYS ============
	// Arrays of the same type can be compared with == (slices cannot!)

	// TODO: Compare two arrays using ==
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{3, 2, 1}
	fmt.Println("\na == b:", a == b) // true
	fmt.Println("a == c:", a == c)   // false

	_ = fmt.Println // remove once used
}
