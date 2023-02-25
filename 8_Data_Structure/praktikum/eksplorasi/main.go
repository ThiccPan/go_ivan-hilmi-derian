package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}
	sum := 0
	for i := 0; i < len(matrix); i++ {
		sum += matrix[i][i]
		sum -= matrix[i][len(matrix)-(i+1)]
	}
	// sum  -= sumLR
	if sum < 0 {
		sum *= -1
	}
	fmt.Println(sum)
}
