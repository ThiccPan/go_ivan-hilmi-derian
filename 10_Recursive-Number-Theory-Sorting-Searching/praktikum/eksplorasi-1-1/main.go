package main

import "fmt"

func MaxSequence(arr []int) {
	// your code here
	max := 0
	sum := 0
	for i := 0; i < len(arr); i++ {
		if sum < 0 {
			sum = 0
		}
		sum += arr[i]
		if sum > max {
			max = sum
		}
	}
	fmt.Println(max)
}

func main() {
	MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})  // 6
	MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})    // 7
	MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})    // 7
	MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})    // 8
    MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})     // 12
}