package main

import (
	"fmt"
	// "strings"
)

func convertBinary(n int) {
	arr := make([]string, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, fmt.Sprintf("%b", i))
	}
	fmt.Println(arr)
}

func main() {
	convertBinary(5)
}
