package main

import (
	"container/heap"
	"fmt"
)

// ============ MIN HEAP BOILERPLATE ============
// Go requires implementing the heap.Interface (5 methods)

// TODO: Define type MinHeap []int

// TODO: Implement Len() int

// TODO: Implement Less(i, j int) bool (return h[i] < h[j] for min heap)

// TODO: Implement Swap(i, j int)

// TODO: Implement Push(x interface{}) - append to the slice
//       func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

// TODO: Implement Pop() interface{} - remove and return last element
//       func (h *MinHeap) Pop() interface{} {
//           old := *h
//           n := len(old)
//           x := old[n-1]
//           *h = old[:n-1]
//           return x
//       }

// ============ MAX HEAP ============
// Same as MinHeap but flip the Less comparison

// TODO: Define type MaxHeap []int
// TODO: Implement all 5 methods (Less should return h[i] > h[j])

func main() {
	// ============ MIN HEAP USAGE ============
	fmt.Println("=== Min Heap ===")

	// TODO: Create h := &MinHeap{}
	// TODO: Initialize with heap.Init(h)
	// TODO: Push values 5, 2, 8, 1, 9 using heap.Push(h, val)
	// TODO: Pop all values and print them (should come out in ascending order)

	fmt.Println()

	// ============ MAX HEAP USAGE ============
	fmt.Println("=== Max Heap ===")

	// TODO: Create mh := &MaxHeap{}
	// TODO: Initialize and push same values: 5, 2, 8, 1, 9
	// TODO: Pop all values and print them (should come out in descending order)

	fmt.Println()

	// ============ PEEK WITHOUT POP ============

	// TODO: Push values 10, 3, 7 into a new min heap
	// TODO: Peek at the minimum without popping: (*h)[0]
	// TODO: Print the peek value

	_ = heap.Init // remove once used
}
