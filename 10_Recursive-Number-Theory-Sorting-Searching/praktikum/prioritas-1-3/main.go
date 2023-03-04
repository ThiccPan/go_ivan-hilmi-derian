package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {
	for i := 2; ; i++ {
		if number < 1 {
			return i-1
		}
		if checkPrime(i) {
			number--
		}
	}
}

func checkPrime(n int) bool {
	if n < 2 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
}