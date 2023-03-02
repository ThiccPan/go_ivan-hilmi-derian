package main

import (
	"fmt"
	"strings"
)

func caesar(offset int, input string) string {
	// your code here
	caesarCip := func(r rune) rune {
		return (((r - 'a') + rune(offset)) % 26) + 'a'
	}

	input = strings.Map(caesarCip, input)
	return input
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // def
	fmt.Println(caesar(10, "alterraacademy"))               //kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
