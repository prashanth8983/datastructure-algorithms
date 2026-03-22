package main

import (
	"context"
	"fmt"
	"time"
)

// ============ CONTEXT WITH TIMEOUT ============

// TODO: Write a function slowOperation(ctx context.Context) error
//       Simulates work with a select:
//       func slowOperation(ctx context.Context) error {
//           select {
//           case <-time.After(2 * time.Second):
//               fmt.Println("operation completed")
//               return nil
//           case <-ctx.Done():
//               fmt.Println("operation cancelled:", ctx.Err())
//               return ctx.Err()
//           }
//       }

// ============ CONTEXT WITH VALUE ============

// TODO: Define a custom key type to avoid collisions:
//       type contextKey string
//       const userKey contextKey = "user"

// TODO: Write a function greetUser(ctx context.Context)
//       that retrieves "user" from context and prints a greeting:
//       func greetUser(ctx context.Context) {
//           user, ok := ctx.Value(userKey).(string)
//           if !ok { fmt.Println("no user in context"); return }
//           fmt.Println("Hello,", user)
//       }

// ============ CONTEXT WITH CANCEL ============

// TODO: Write a function worker(ctx context.Context, id int)
//       that loops printing "working" until ctx is cancelled:
//       func worker(ctx context.Context, id int) {
//           for {
//               select {
//               case <-ctx.Done():
//                   fmt.Printf("worker %d stopped\n", id)
//                   return
//               default:
//                   fmt.Printf("worker %d working\n", id)
//                   time.Sleep(200 * time.Millisecond)
//               }
//           }
//       }

func main() {
	// ============ TIMEOUT ============
	fmt.Println("=== Context with Timeout ===")

	// TODO: Create a context with 500ms timeout:
	//       ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	//       defer cancel()
	// TODO: Call slowOperation(ctx) — should be cancelled (500ms < 2s)

	// ============ VALUE ============
	fmt.Println("\n=== Context with Value ===")

	// TODO: Create a context with a value:
	//       ctx := context.WithValue(context.Background(), userKey, "Alice")
	// TODO: Call greetUser(ctx)
	// TODO: Call greetUser(context.Background()) — no user set

	// ============ CANCEL ============
	fmt.Println("\n=== Context with Cancel ===")

	// TODO: Create a cancellable context:
	//       ctx, cancel := context.WithCancel(context.Background())
	// TODO: Launch worker(ctx, 1) as a goroutine
	// TODO: Sleep 1 second, then call cancel()
	// TODO: Sleep a bit more to see the worker stop

	_ = context.Background   // remove once used
	_ = time.Second          // remove once used
	_ = fmt.Println          // remove once used
}
