package main

import "fmt"

// jika n % bilangan prima 1 digit == 0, maka bukan prima
func primeNumber(n int) bool {
	// bilangan prima 1 digit
	if n == 2 || n == 3 || n == 5 || n == 7 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	if n%3 == 0 {
		return false
	}

	if n%5 == 0 {
		return false
	}

	if n&7 == 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println("1023 adalah prima: ", primeNumber(1023))
	fmt.Println("1021 adalah prima: ", primeNumber(1021))
	fmt.Println("248 adalah prima: ", primeNumber(248))
}
