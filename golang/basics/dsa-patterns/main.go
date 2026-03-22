package main

import "fmt"

func main() {
	// ============ STACK ============
	fmt.Println("=== Stack ===")

	// TODO: Create a stack using []int{}
	// TODO: Push values 1, 2, 3
	// TODO: Peek at top: stack[len(stack)-1]
	// TODO: Pop from top: stack = stack[:len(stack)-1]
	// TODO: Print stack after each operation

	// ============ QUEUE ============
	fmt.Println("\n=== Queue ===")

	// TODO: Create a queue using []int{}
	// TODO: Enqueue values 1, 2, 3
	// TODO: Dequeue from front: front := queue[0]; queue = queue[1:]
	// TODO: Print queue after each operation

	// ============ BFS ============
	fmt.Println("\n=== BFS ===")

	// Graph as adjacency list
	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {5},
		4: {},
		5: {},
	}

	// TODO: Implement BFS starting from node 1
	//       - Create queue := []int{1}
	//       - Create visited := map[int]bool{1: true}
	//       - While queue is not empty:
	//         - Dequeue front node
	//         - Print it
	//         - For each neighbor, if not visited, mark visited and enqueue

	_ = graph

	// ============ BINARY SEARCH ============
	fmt.Println("\n=== Binary Search ===")

	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 7

	// TODO: Implement binary search
	//       lo, hi := 0, len(arr)-1
	//       Loop while lo <= hi
	//       mid := lo + (hi-lo)/2  (avoids overflow)
	//       If found, print index; if not found, print -1

	_ = arr
	_ = target

	// ============ BACKTRACKING ============
	fmt.Println("\n=== Backtracking (Subsets) ===")

	nums := []int{1, 2, 3}

	// TODO: Generate all subsets of nums using backtracking
	//       var result [][]int
	//       var backtrack func(start int, path []int)
	//       backtrack = func(start int, path []int) {
	//           // IMPORTANT: make a copy before appending to result
	//           tmp := make([]int, len(path))
	//           copy(tmp, path)
	//           result = append(result, tmp)
	//
	//           for i := start; i < len(nums); i++ {
	//               path = append(path, nums[i])
	//               backtrack(i+1, path)
	//               path = path[:len(path)-1]  // undo choice
	//           }
	//       }
	//       backtrack(0, []int{})

	_ = nums

	_ = fmt.Println // remove once used
}
