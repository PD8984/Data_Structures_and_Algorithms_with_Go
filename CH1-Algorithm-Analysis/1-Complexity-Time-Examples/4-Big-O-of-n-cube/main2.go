// Time Complexity: O(n^3)

package main

import (
	"fmt"
)

func funEleven(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := j + 1; k < n; k++ {
				m += 1
			}
		}
	}
	return m
}

func main() {
	fmt.Println("N = 100, Number of instruction 0(n^3) ::", funEleven(100))
}
