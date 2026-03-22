package main

import "fmt"

// ============ DEFER ============
// Defer schedules a function to run when the surrounding function returns
// Deferred calls execute in LIFO (stack) order

func main() {
	fmt.Println("=== Defer Basics ===")

	// TODO: Add three deferred prints and a normal print:
	//       defer fmt.Println("deferred 1")
	//       defer fmt.Println("deferred 2")
	//       defer fmt.Println("deferred 3")
	//       fmt.Println("normal")
	// Expected output: normal, deferred 3, deferred 2, deferred 1

	// ============ DEFER FOR CLEANUP ============
	fmt.Println("\n=== Defer for Cleanup ===")

	// TODO: Write a function readFile() that demonstrates defer for cleanup:
	//       func readFile() {
	//           fmt.Println("opening file")
	//           defer fmt.Println("closing file")  // always runs, even on error
	//           fmt.Println("reading file")
	//           // even if panic happens here, defer still runs
	//       }
	// TODO: Call readFile()

	// ============ DEFER WITH LOOPS ============
	fmt.Println("\n=== Defer in Loop ===")

	// TODO: Loop i from 0 to 4 and defer printing each value
	//       for i := 0; i < 5; i++ {
	//           defer fmt.Print(i, " ")
	//       }
	//       fmt.Println("before deferred:")
	// Expected: prints 4 3 2 1 0 (LIFO order)

	// ============ PANIC ============
	fmt.Println("\n=== Panic ===")

	// TODO: Write a function mustPositive(n int) that panics if n < 0:
	//       func mustPositive(n int) {
	//           if n < 0 { panic("negative number!") }
	//           fmt.Println("n is", n)
	//       }

	// ============ RECOVER ============
	fmt.Println("\n=== Recover ===")

	// TODO: Write a function safeCall(f func()) that recovers from panics:
	//       func safeCall(f func()) {
	//           defer func() {
	//               if r := recover(); r != nil {
	//                   fmt.Println("Recovered from panic:", r)
	//               }
	//           }()
	//           f()
	//       }

	// TODO: Call safeCall with a function that panics:
	//       safeCall(func() { panic("oops!") })
	//       fmt.Println("program continues after panic")

	// TODO: Call safeCall with mustPositive(-1)
	//       safeCall(func() { mustPositive(-1) })

	_ = fmt.Println // remove once used
}
