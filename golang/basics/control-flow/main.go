package main

import "fmt"

func compute() int {
	return 8
}

func main() {
	// ============ IF STATEMENTS ============
	// Go: no parens around condition, init statement allowed

	// TODO: Write an if statement with init statement: x := compute()
	//       If x > 10 print "big", else if x > 5 print "medium", else print "small"

	if x := compute(); x>10{
		fmt.Println("big")
	}else if x>5{
		fmt.Println("medium")
	}else{
		fmt.Println("small")
	}

	// ============ FOR LOOPS (the only loop in Go) ============
	n := 5

	// TODO: Classic for loop - print numbers 0 to n-1
	fmt.Print("Classic: ")

	for i:=0; i<n; i++ {
		fmt.Print(i," ")
	}

	fmt.Println()

	// TODO: While-style for loop - print numbers 0 to n-1 using only a condition
	fmt.Print("While: ")
	i := 0
	for i<n{
		fmt.Print(i, " ");
		i++
	}

	fmt.Println()

	// TODO: Range over a slice - print each value of []int{10, 20, 30}
	fmt.Print("Range slice: ")

	slice := []int{10,20,30}

	for _ ,v := range slice{
		fmt.Print(v," ")
	}

	fmt.Println()

	// TODO: Range over a map - print each key-value pair
	//       Use map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Print("Range map: ")

	myMap := map[string]int{"a":1, "b":2, "c":3 }

	for k,v := range myMap{
		fmt.Print("\n",k,":",v)
	}

	fmt.Println()

	// TODO: Range over a string "hello" - print each rune (character)
	fmt.Print("Range string: ")

	s := []rune("hello")

	for _, v := range s{
		fmt.Print(string(v))
	}

	fmt.Println()

	// ============ SWITCH ============
	day := 3

	// TODO: Write a switch on day: case 1 -> "Mon", case 2 -> "Tue", case 3 -> "Wed",
	//       cases 4,5 -> "Thu/Fri", default -> "Weekend"
	//       (no break needed in Go, no fallthrough by default)

	switch day {

		case 1: fmt.Print("Mon")
		case 2: fmt.Print("Tue")
		case 3: fmt.Print("Wed")
		case 4,5: fmt.Print("Thu/Fri")
		default: fmt.Print("Weekend")
	}

	// TODO: Write a switch with NO condition (acts like if-else chain)
	//       If day > 5 print "Weekend", else if day > 3 print "Late week", else print "Early week"
	switch {
	case day> 5: fmt.Print("weekend")
	case day>3: fmt.Print("late week")
	default:
			fmt.Print("early week")
	}
	_ = fmt.Println // remove once used
}
