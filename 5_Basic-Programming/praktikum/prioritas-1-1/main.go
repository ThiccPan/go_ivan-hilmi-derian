package main

import (
	"fmt"
)

func main() {
	// deklarasi variable yang dibutuhkan
	var sisiA, sisiB, tinggi int

	convInput(&sisiA, "input sisi a: ")

	convInput(&sisiB, "input sisi b: ")

	convInput(&tinggi, "input tinggi: ")

	luas := (sisiA + sisiB) / 2 * tinggi
	fmt.Println(luas)

}

// ambil input dari stdin(string), lalu convert jadi int
func convInput(trapesiumVar *int, msg string) {
	fmt.Println(msg)
	fmt.Scanf("%v", trapesiumVar)

}