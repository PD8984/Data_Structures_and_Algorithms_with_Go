// In each iteration, the value of i is halved. Each iteration problem space is divide into half.
// Time Complexity: O(log(n))

package main

import (
	"fmt"
)

func funFive(n int) int {
	m := 0
	i := 0
	for i > 0 {
		m += 1
		i = i / 2
	}
	return m
}

func main() {
	fmt.Println("N = 100, Number of instruction O(log(n)) ::", funFive(100))
}
