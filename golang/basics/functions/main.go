package main

import (
	"fmt"
	"sort"
)

// ============ BASIC FUNCTION ============

// TODO: Write a function add(a, b int) int that returns a + b

func add(a , b int) int{
	return a+b
}

// ============ MULTIPLE RETURNS ============

// TODO: Write a function divmod(a, b int) (int, int) that returns quotient and remainder

// ============ VARIADIC FUNCTION ============

func divmod(a,b int) (int, int){
	quo := a/b
	rem := a%b
	return quo, rem
}

// TODO: Write a function sum(nums ...int) int that returns the sum of all nums

func sum(nums ...int) int{
	sum := 0
	for _,v := range nums{
		sum += v
	}
	return sum
}

func main() {
	// TODO: Call add(3, 4) and print result

	result := add(3,4)
	fmt.Print(result)

	// TODO: Call divmod(10, 3), capture both values (q, r), print them

	q,r := divmod(10,3)
	fmt.Print(q," ",r)

	// TODO: Call sum(1, 2, 3, 4, 5) and print result

	res := sum(1,2,3,4,5,6)
	fmt.Print(res)

	// TODO: Call sum with a slice using spread: nums := []int{1,2,3}; sum(nums...)

	nums := []int{1,2,3}; sum(nums...)
	// ============ CLOSURE (critical for DSA recursion) ============

	// Simulating a tree
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 3},
	}

	// TODO: Declare a closure variable: var dfs func(node *TreeNode)
	// TODO: Implement dfs to print node values in preorder (root, left, right)
	//       Pattern:
	//       dfs = func(node *TreeNode) {
	//           if node == nil { return }
	//           fmt.Print(node.Val, " ")
	//           dfs(node.Left)
	//           dfs(node.Right)
	//       }

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode){
		if node == nil {return}
		fmt.Print(node.Val," ")
		dfs(node.Left)
		dfs(node.Right)
	}

	

	fmt.Print("Preorder: ")
	// TODO: Call dfs(root)
	dfs(root)
	fmt.Println()

	// ============ ANONYMOUS FUNCTION FOR SORTING ============
	arr := []int{5, 2, 8, 1, 9}

	// TODO: Sort arr in descending order using sort.Slice with anonymous function
	//       sort.Slice(arr, func(i, j int) bool { return arr[i] > arr[j] })


	sort.Slice(arr, func(i,j int) bool {return arr[i]> arr[j]})

	fmt.Println("Sorted desc:", arr)

	_ = sort.Slice // remove once used
}
