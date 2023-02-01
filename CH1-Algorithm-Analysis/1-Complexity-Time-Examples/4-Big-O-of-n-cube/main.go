// Outer loop will run for n number of iterations.
// In each iteration of the outer loop, inner loop will run for n iteration of its own.
// Three nested loops running number of three times.
// Time Complexity: O(n^3)

package main

import (
	"fmt"
)

func funSix(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				m += 1
			}
		}
	}
	return m
}

func main() {
	fmt.Println("N = 100, Number of instruction O(log(n)) ::", funSix(100))
}
