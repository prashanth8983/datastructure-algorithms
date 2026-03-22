package main

import "fmt"

// ============ STRUCT DEFINITIONS ============

// TODO: Define a TreeNode struct with fields:
//       Val   int
//       Left  *TreeNode
//       Right *TreeNode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


// TODO: Define a Pair struct with fields x, y int

type Pair struct {
	X,Y int
}

// TODO: Define a method IsLeaf on *TreeNode that returns bool
//       A leaf has no left or right children
// func (n *TreeNode) IsLeaf() bool { ... }

func (n *TreeNode) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func main() {
	// ============ CREATING STRUCTS ============

	// TODO: Create a TreeNode pointer with Val=5 using &TreeNode{Val: 5}
	var root *TreeNode
	root = &TreeNode{Val: 5}
	// TODO: Set root.Left to a new TreeNode with Val=3
	root.Left = &TreeNode{Val: 3}
	// TODO: Set root.Right to a new TreeNode with Val=7
	root.Right = &TreeNode{Val: 7}

	// TODO: Print whether root is a leaf
	if root.IsLeaf(){
		fmt.Println("True")
	}else{
		fmt.Println("False")
	}
	// TODO: Print whether root.Left is a leaf
	if root.Left != nil{

		if root.Left.IsLeaf() {
			fmt.Println("True")
		}else{
			fmt.Println("False")
		}
	}

	fmt.Println("Root:", root.Val)
	fmt.Println("Root is leaf:", root.IsLeaf())
	fmt.Println("Left is leaf:", root.Left.IsLeaf())

	// TODO: Create a Pair with x=1, y=2

	p := Pair{X:1, Y:2}

	fmt.Println("Pair:", p)

	// ============ STRUCT AS MAP KEY ============
	// Structs can be map keys if all fields are comparable

	// TODO: Create a map visited := map[Pair]bool{}

	visited := map[Pair]bool{}
	// TODO: Mark Pair{0, 0} and Pair{1, 1} as visited
	visited[Pair{0,0}] = true
	visited[Pair{1,1}] = true

	// TODO: Check if Pair{0, 0} is visited

	_, ok := visited[Pair{0,0}]

	if ok{
		fmt.Println("Pair exists")
	}else{
		fmt.Println("Pair doesn't exists")
	}

	fmt.Println("Visited (0,0):", visited[Pair{0, 0}])
	fmt.Println("Visited (2,2):", visited[Pair{2, 2}])
}
