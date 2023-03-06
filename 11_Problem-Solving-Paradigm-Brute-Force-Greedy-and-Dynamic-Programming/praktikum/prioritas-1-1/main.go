package main

import (
	"fmt"
	// "strings"
)

func convertRomawi(n int) {
	arr := make([]string, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, fmt.Sprintf("%b", i))
	}
	fmt.Println(arr)
}

func main() {
	convertRomawi(5)
}
