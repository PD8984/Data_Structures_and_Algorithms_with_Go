// For nested loops look for inner loop iterations.
// Time complexity will be calculated by looking into inner loop.
// First, it will run for n number of time then n/2 ans so on. (n + n/2 + n/4 + n/16)

package main

import (
	"fmt"
)

func funNine(n int) int {
	m := 0
	for i := n; i > 0; i /= 2 {
		for j := 0; j < i; j++ {
			m += 1
		}
	}
	return m
}

func main() {
	fmt.Println("N = 100, Number of instruction O(n) ::", funNine(100))
}
