package main

import (
	"fmt"
)


func main() {
	fmt.Println("masukkan nilai: ")

	var nilai int8
	fmt.Scanf("%v", &nilai)

	switch {
	case nilai > 100:
		fmt.Println("Invalid input")
	case nilai >= 80:
		fmt.Println("A")
	case nilai >= 65:
		fmt.Println("B")
	case nilai >= 50:
		fmt.Println("C")
	case nilai >= 35:
		fmt.Println("D")
	case nilai >= 0:
		fmt.Println("E")
	default:
		fmt.Println("Invalid input")
	}
	

}