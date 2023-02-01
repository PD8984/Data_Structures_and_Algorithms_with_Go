// Two group of loops are in consecutive so their complexity will add up for up to form the final complexity of the program.:: Time Complexity : O(n^2) + O(n^2) = O(n^2)

package main

import (
	"fmt"
)

func funSeven(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m += 1
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m += 1
		}
	}
	return m
}

func main() {
	fmt.Println("N = 100, Number of instructions 0(n^2) :: ", funSeven(100))
}
