// Time Complexity: O(n²), two nested for loop take quadratic time.

package main

import (
  "fmt"
)

func funTwo (n int) int {
  m := 0
  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      m += 1
    }
  }
  return m
}

func main() {
  fmt.Println("N = 100, Number of instructions 0(n^2) :: ", funTwo(100))
}
