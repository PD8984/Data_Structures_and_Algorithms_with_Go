// Time Complexity: O(n), single for loop takes linear time

package main

import (
  "fmt"
)

func funOne(n int) int {
  m := 0
  for i := 0; i < n; i++ {
    m += 1
  }
  return m
}

func main() {
  fmt.Println("N = 100, Number of instructions O(n) :: ", funOne(100))
}
