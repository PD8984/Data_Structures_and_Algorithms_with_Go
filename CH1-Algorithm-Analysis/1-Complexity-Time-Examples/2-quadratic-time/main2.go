// Time Complexity: o(N + (N-1) + (N-2) + ...) == O(N(N+1/2)) == O(N^2)

package main

import (
    "fmt"
)

func funThree(n int) int {
    m := 0
    for i := 0; i < n; i++ {
        for j := 0; j < i; j++ {
            m += 1
        }
    }
    return m
}

func main() {
    fmt.Println("N = 100, Number of instructions for O(n^2) ::", funThree(100))
}
