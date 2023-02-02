// The inner loop with run for 1, 2, 4, 8 ... n times in sucessive iteration of the outer loop
// Time Complexity: T(n) = O(1 + 2 + 4 + ... + n/2+n) = O(n)

package main

import (
	"fmt"
)

func funThirteen(n int) int {
	m := 0
	for i := 1; i <= n; i *= 2 {
		for j := 0; j <= i; j++ {
			m += 1
		}
	}
	return m
}

func main() {
	fmt.Println("N = 100, Number of instruction 0(n) ::", funThirteen(100))
}
