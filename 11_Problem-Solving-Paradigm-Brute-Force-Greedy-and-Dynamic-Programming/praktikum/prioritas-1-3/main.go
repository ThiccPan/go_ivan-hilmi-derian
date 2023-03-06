package main

import "fmt"

func fibonacci(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(fibonacci(0)) // 0
	fmt.Println(fibonacci(1)) // 1
	fmt.Println(fibonacci(2)) // 1
	fmt.Println(fibonacci(3)) // 2
	fmt.Println(fibonacci(5)) // 5
	fmt.Println(fibonacci(6)) // 8
	fmt.Println(fibonacci(7)) // 13
	fmt.Println(fibonacci(8)) // 21
	fmt.Println(fibonacci(9)) // 34
	fmt.Println(fibonacci(10)) // 55
}