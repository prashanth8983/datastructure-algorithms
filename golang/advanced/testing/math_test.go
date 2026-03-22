package main

import "testing"

// ============ BASIC TEST ============

// TODO: Write a test for Add:
//       func TestAdd(t *testing.T) {
//           result := Add(2, 3)
//           if result != 5 {
//               t.Errorf("Add(2, 3) = %d; want 5", result)
//           }
//       }

// ============ TABLE-DRIVEN TESTS ============
// This is the idiomatic Go testing pattern

// TODO: Write a table-driven test for Max:
//       func TestMax(t *testing.T) {
//           tests := []struct {
//               name     string
//               a, b     int
//               expected int
//           }{
//               {"both positive", 3, 5, 5},
//               {"both negative", -3, -5, -3},
//               {"equal", 4, 4, 4},
//               {"mixed", -1, 1, 1},
//           }
//           for _, tt := range tests {
//               t.Run(tt.name, func(t *testing.T) {
//                   result := Max(tt.a, tt.b)
//                   if result != tt.expected {
//                       t.Errorf("Max(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
//                   }
//               })
//           }
//       }

// ============ TEST FOR PALINDROME ============

// TODO: Write a table-driven test for IsPalindrome with cases:
//       "racecar" -> true, "hello" -> false, "" -> true, "a" -> true, "ab" -> false

// ============ BENCHMARK ============

// TODO: Write a benchmark for Fibonacci:
//       func BenchmarkFibonacci(b *testing.B) {
//           for i := 0; i < b.N; i++ {
//               Fibonacci(20)
//           }
//       }
// Run with: go test -bench=.
