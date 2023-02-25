package main

import "fmt"

func PairSum(arr []int, target int) []int {
	mapNum := make(map[int]int)
	pair := []int{}
	for i, arrVal := range arr {
		mapVal, isPresent := mapNum[target-arrVal]
		
		if isPresent {
			pair = append(pair, mapVal, i)
			return pair
		} else {
			mapNum[arrVal] = i
		}
	}

	return pair
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
}
