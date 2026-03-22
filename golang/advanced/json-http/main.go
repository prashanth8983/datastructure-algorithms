package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ============ JSON STRUCT TAGS ============

// TODO: Define a struct User with JSON tags:
//       type User struct {
//           Name  string `json:"name"`
//           Email string `json:"email"`
//           Age   int    `json:"age,omitempty"`  // omitted if zero
//       }

// ============ MARSHAL (struct -> JSON) ============

// TODO: Write a function toJSON(u User) string
//       data, err := json.Marshal(u)
//       if err != nil { return "" }
//       return string(data)

// For pretty print: json.MarshalIndent(u, "", "  ")

// ============ UNMARSHAL (JSON -> struct) ============

// TODO: Write a function fromJSON(s string) (User, error)
//       var u User
//       err := json.Unmarshal([]byte(s), &u)
//       return u, err

// ============ SIMPLE HTTP SERVER ============

// TODO: Write a handler function:
//       func helloHandler(w http.ResponseWriter, r *http.Request) {
//           fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
//       }

// TODO: Write a JSON API handler:
//       func userHandler(w http.ResponseWriter, r *http.Request) {
//           user := User{Name: "Alice", Email: "alice@example.com", Age: 30}
//           w.Header().Set("Content-Type", "application/json")
//           json.NewEncoder(w).Encode(user)
//       }

func main() {
	// ============ JSON ============
	fmt.Println("=== JSON ===")

	// TODO: Create a User and convert to JSON, print it
	// TODO: Parse a JSON string back to User, print it
	//       jsonStr := `{"name":"Bob","email":"bob@example.com","age":25}`

	// ============ HTTP SERVER ============
	fmt.Println("\n=== HTTP Server ===")
	fmt.Println("Uncomment the lines below to start the server")

	// TODO: Register handlers and start server:
	//       http.HandleFunc("/hello/", helloHandler)
	//       http.HandleFunc("/api/user", userHandler)
	//       fmt.Println("Server starting on :8080")
	//       http.ListenAndServe(":8080", nil)

	// Test with:
	//   curl http://localhost:8080/hello/World
	//   curl http://localhost:8080/api/user

	_ = json.Marshal         // remove once used
	_ = http.ListenAndServe  // remove once used
	_ = fmt.Println          // remove once used
}
