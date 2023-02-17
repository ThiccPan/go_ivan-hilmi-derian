package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("masukkan angka")

	fmt.Scanf("%v", &n)

	if n%2 == 0 {
		fmt.Println("genap")
	} else {
		fmt.Println("ganjil")
	}
}
