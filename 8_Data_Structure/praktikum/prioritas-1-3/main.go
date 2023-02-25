package main

import (
	"fmt"
)

// O(N) time complexity, O(1) space Complexity
func munculSekali(angka string) []int {
	// index num menggambarkan char angka, value menggambarkan jumlah muncul
	num := [10]uint8{}
	
	// mencatat jumlah muncul angka dalam string, kalo < 2 bakal ditambah
	for _, v := range angka {
		if num[uint8(v)-48] < 2 {
			num[uint8(v)-48]++
		}
	}
	// O(N) time, O(1) space, iterasi loop sebanyak panjang string "angka"

	uNum := []int{}
	// cek banyak muncul bilangan, kalo nggak muncul tepat 1x nggak masuk arr uNum
	for i, v := range num {
		if v == 1 {
			uNum = append(uNum, i)
		}
	}
	// O(1) time & space, range maks adalah 10 (digit 0 - 9)
	return uNum
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
