package main

import (
  "fmt"
)

func bubbleSort(arr []int, comp func(int, int) bool) {
  size := len(arr)
  // Outer loop represents the number of swaps that are done for comparison
  for i := 0; i < (size-1); i++ {
    // Inner loop is used to do the comparison of data
    for j := 0; j < (size-i-1); j++ {
      if comp(arr[j], arr[j+1]) {
        //swapping 
        arr[j+1], arr[j] = arr[j], arr[j+1]
      }
    }
  }
}

func more(value1 int, value2 int) bool {
  return value1 > value2
}

func main() {
  data := []int{9,1,8,2,7,3,6,4,5}
  bubbleSort(data, more)
  fmt.Println(data)
}
