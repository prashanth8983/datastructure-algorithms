package main

import (
	"fmt"
	"sync"
	"time"
)

// ============ BASIC GOROUTINE ============

func sayHello(name string) {
	fmt.Println("Hello from", name)
}

func main() {
	// ============ LAUNCHING GOROUTINES ============
	fmt.Println("=== Basic Goroutines ===")

	// TODO: Launch sayHello("goroutine") as a goroutine using go keyword
	// TODO: Call sayHello("main") normally
	// TODO: Add time.Sleep(100 * time.Millisecond) to wait for goroutine
	//       (this is a bad pattern — we'll fix it with WaitGroup below)

	// ============ WAITGROUP ============
	fmt.Println("\n=== WaitGroup ===")

	// TODO: Create a sync.WaitGroup
	// TODO: Launch 5 goroutines, each printing their index
	//       var wg sync.WaitGroup
	//       for i := 0; i < 5; i++ {
	//           wg.Add(1)
	//           go func(id int) {
	//               defer wg.Done()
	//               fmt.Println("Worker", id)
	//           }(i)     // pass i as argument to avoid closure bug!
	//       }
	//       wg.Wait()

	// ============ CHANNELS ============
	fmt.Println("\n=== Channels ===")

	// TODO: Create an unbuffered channel: ch := make(chan string)
	// TODO: Launch a goroutine that sends "hello" to the channel
	//       go func() { ch <- "hello" }()
	// TODO: Receive from channel and print: msg := <-ch

	// ============ BUFFERED CHANNEL ============
	fmt.Println("\n=== Buffered Channel ===")

	// TODO: Create a buffered channel with capacity 3: ch := make(chan int, 3)
	// TODO: Send 1, 2, 3 without blocking (buffer has space)
	// TODO: Receive and print all 3 values

	// ============ CHANNEL WITH RANGE ============
	fmt.Println("\n=== Channel Range ===")

	// TODO: Create a channel, launch a goroutine that sends 1..5 then closes
	//       ch := make(chan int)
	//       go func() {
	//           for i := 1; i <= 5; i++ { ch <- i }
	//           close(ch)   // MUST close or range will deadlock
	//       }()
	// TODO: Range over channel and print each value
	//       for v := range ch { fmt.Print(v, " ") }

	// ============ SELECT ============
	fmt.Println("\n=== Select ===")

	// TODO: Create two channels ch1, ch2
	// TODO: Send values on each from goroutines
	// TODO: Use select to receive from whichever is ready first:
	//       select {
	//       case msg1 := <-ch1:
	//           fmt.Println("from ch1:", msg1)
	//       case msg2 := <-ch2:
	//           fmt.Println("from ch2:", msg2)
	//       }

	// ============ MUTEX ============
	fmt.Println("\n=== Mutex ===")

	// TODO: Demonstrate safe concurrent counter using sync.Mutex
	//       var mu sync.Mutex
	//       counter := 0
	//       var wg2 sync.WaitGroup
	//       for i := 0; i < 1000; i++ {
	//           wg2.Add(1)
	//           go func() {
	//               defer wg2.Done()
	//               mu.Lock()
	//               counter++
	//               mu.Unlock()
	//           }()
	//       }
	//       wg2.Wait()
	//       fmt.Println("Counter:", counter)  // always 1000

	_ = sync.WaitGroup{}         // remove once used
	_ = time.Millisecond         // remove once used
	_ = fmt.Println              // remove once used
}
