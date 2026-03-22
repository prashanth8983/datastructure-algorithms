package main

import "fmt"

func main() {
	// ============ MAP BASICS ============

	// TODO: Create an empty map m with string keys and int values


	m := map[string]int{}

	// TODO: Set m["apple"] = 5
	m["apple"] = 5
	// TODO: Get the value for "apple" and print it

	fmt.Print(m["apple"])

	// TODO: Use the two-value form (val, ok) to check if "banana" exists

	val, ok := m["banana"]

	if ok{
		fmt.Println(val)
	}

	// TODO: Delete "apple" from the map
	delete(m, "apple")
	

	// TODO: Print the length of the map

	fmt.Println(len(m))

	// ============ ITERATE OVER MAP ============

	scores := map[string]int{"alice": 90, "bob": 85, "carol": 95}

	// TODO: Range over scores and print each key-value pair
	//       (note: iteration order is random in Go!)

	for k,v := range scores{
		fmt.Println(k," : ", v)
	}



	// ============ SET PATTERN ============

	// TODO: Create a "set" using map[int]bool{} called seen

	seen := map[int]bool{}
	// TODO: Add values 1, 2, 3 to the set
	seen[1] = true
	seen[2] = true
	seen[3] = true
	// TODO: Check if 2 is in the set and print result

	if seen[2]{
		fmt.Println(seen[2])
	}

	// ============ DEFAULT MAP (like Python defaultdict) ============

	// TODO: Create a frequency counter freq := map[string]int{}

	freq := map[string]int{}

	

	// TODO: Increment freq["a"] three times (works because zero value of int is 0)
	freq["a"]++
	freq["a"]++
	freq["a"]++
	// TODO: Increment freq["b"] once

	freq["b"]++
	// TODO: Print freq

	for k,v := range freq{
		fmt.Println(k,": ",v)
	}

	// ============ MAP OF SLICES (adjacency list) ============

	// TODO: Create a graph := map[int][]int{}

	graph := map[int][]int{}
	// TODO: Add edge 1->2, 1->3, 2->3
	//       Use: graph[1] = append(graph[1], 2) pattern

	graph[1] = append(graph[1],2)
	graph[1] = append(graph[1],3)
	graph[2] = append(graph[2],3)

	// TODO: Print the graph

	for k := range graph{
		fmt.Print(k,"->")
		for _,n := range graph[k]{
			fmt.Print(n," ")
		}
		fmt.Println()
	}
}
