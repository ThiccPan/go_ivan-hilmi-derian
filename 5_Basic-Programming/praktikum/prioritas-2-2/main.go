package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

// loop i sampai angka, jika angka habis dibagi i print i
func main() {
	fmt.Println("angka: ")
	scanner.Scan()

	var angka int
	fmt.Sscan(scanner.Text(), &angka)
	for i := 1; i <= angka; i++ {
		if angka % i == 0 {
			fmt.Println(i)
		}
	}
}