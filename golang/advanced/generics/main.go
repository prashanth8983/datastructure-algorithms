package main

import "fmt"

// ============ GENERIC FUNCTION ============

// TODO: Write a generic function Min that works with any ordered type:
//       func Min[T int | float64 | string](a, b T) T {
//           if a < b { return a }
//           return b
//       }

// ============ TYPE CONSTRAINTS ============

// TODO: Define a constraint interface:
//       type Number interface {
//           int | int32 | int64 | float32 | float64
//       }

// TODO: Write a generic Sum function:
//       func Sum[T Number](nums []T) T {
//           var total T
//           for _, n := range nums {
//               total += n
//           }
//           return total
//       }

// ============ GENERIC STRUCT ============

// TODO: Define a generic Stack:
//       type Stack[T any] struct {
//           items []T
//       }
//       func (s *Stack[T]) Push(item T) { s.items = append(s.items, item) }
//       func (s *Stack[T]) Pop() (T, bool) {
//           var zero T
//           if len(s.items) == 0 { return zero, false }
//           item := s.items[len(s.items)-1]
//           s.items = s.items[:len(s.items)-1]
//           return item, true
//       }
//       func (s *Stack[T]) Peek() (T, bool) {
//           var zero T
//           if len(s.items) == 0 { return zero, false }
//           return s.items[len(s.items)-1], true
//       }
//       func (s *Stack[T]) Len() int { return len(s.items) }

// ============ GENERIC MAP/FILTER ============

// TODO: Write a generic Map function:
//       func Map[T any, U any](slice []T, f func(T) U) []U {
//           result := make([]U, len(slice))
//           for i, v := range slice {
//               result[i] = f(v)
//           }
//           return result
//       }

// TODO: Write a generic Filter function:
//       func Filter[T any](slice []T, f func(T) bool) []T {
//           var result []T
//           for _, v := range slice {
//               if f(v) { result = append(result, v) }
//           }
//           return result
//       }

func main() {
	// TODO: Test Min with ints, floats, and strings

	// TODO: Test Sum with []int{1, 2, 3, 4, 5} and []float64{1.1, 2.2, 3.3}

	// TODO: Create a Stack[int], push 1, 2, 3, pop and print

	// TODO: Create a Stack[string], push "a", "b", "c", peek and print

	// TODO: Use Map to double each number in []int{1, 2, 3}

	// TODO: Use Filter to keep only even numbers from []int{1, 2, 3, 4, 5, 6}

	_ = fmt.Println // remove once used
}
