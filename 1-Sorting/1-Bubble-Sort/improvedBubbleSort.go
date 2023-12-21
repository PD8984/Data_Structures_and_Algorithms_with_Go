package main

import (
  "fmt"
)

func improvedBubbleSort(arr []int, comp func(int, int) bool) {
  size := len(arr)
  swapped := 1
  for i := 0; i < (size - 1) && swapped == 1; i++ {
    swapped = 0
    for j := 0; j < size-i-1; j++ {
      if comp(arr[j], arr[j+1]) {
        arr[j+1], arr[j] = arr[j], arr[j+1]
        swapped = 1
      }
    }
  }
}

func more(value1 int, value2 int) bool {
  return value1 > value2
}

func main() {
  data := []int{9,1,8,2,7,3,6,4,5}
  improvedBubbleSort(data, more)
  fmt.Println(data)
}
