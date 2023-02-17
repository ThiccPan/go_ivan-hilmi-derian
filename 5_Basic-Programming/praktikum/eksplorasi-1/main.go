package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	sentence := scanner.Text()
	isPalindrome := true

	fmt.Print("captured: ")

	var opposite byte
	for i, v := range sentence {
		opposite = sentence[len(sentence)-(i+1)]

		if byte(v) != opposite {
			isPalindrome = false
		}
		fmt.Print(string(opposite))
	}

	if isPalindrome {
		fmt.Println("\npalindrome")
	} else {
		fmt.Println("\nbukan palidrome")
	}
}
