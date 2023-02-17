package main

import "fmt"

func main() {
	fmt.Println(0)
	for i := 1; i < 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Print("fizzbuzz")
		case i%3 == 0:
			fmt.Print("fizz")
		case i%5 == 0:
			fmt.Print("buzz")
		default:
			fmt.Print(i)
		}
		fmt.Println()
	}
}
