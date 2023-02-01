// Time Complexity: (n^(3/2))

package main

import (
	"fmt"
	"math"
)

func funEight(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		sq := math.Sqrt(float64(n))
		for j := 0; j < int(sq); j++ {
			m += 1
		}
	}
	return m
}

func main() {
	fmt.Println("N = 100, Number of instructions O(n^(3/2)) ::", funEight(100))
}
