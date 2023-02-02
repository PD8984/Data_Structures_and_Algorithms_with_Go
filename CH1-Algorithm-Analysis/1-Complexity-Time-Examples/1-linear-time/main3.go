// Time Complexity: O(n)

package main

import (
	"fmt"
)

func funTwelve(n int) int {
	j := 0
	m := 0
	for i := 0; i < n; i++ {
		for ; j < n; j++ { //j value is not reset at eat iteration
			m += 1
		}
	}
	return m
}

func main() {
	fmt.Println("N = 100, Number of instructions O(n) ::", funTwelve(100))
}
