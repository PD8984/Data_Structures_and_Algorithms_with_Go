// O(N + (N-1) + (N-2) + ...) = O(N(N+1)/2)
// Time Complexity: O(n^2)

package main

import "fmt"

func funTen(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			m += 1
		}
	}
	return m
}

func main() {
	fmt.Println("N = 100, Number of instructions O(n^2) ::", funTen(100))
}
