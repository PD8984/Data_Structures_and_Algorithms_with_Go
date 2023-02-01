// In each iteration, the value of i is doubled. Each iteration problem space is divide into half.
// Time Complexity: O(log(n))

package main

import (
	"fmt"
)

func funFour(n int) int {
	m := 0
	i := 1
	for i < n {
		m += 1
		i = i * 2
	}
	return m
}

func main() {
	fmt.Println("N = 100, Number of instruction O(log(n)) ::", funFour(100))
}
