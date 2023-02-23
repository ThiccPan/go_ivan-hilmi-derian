package main

import "fmt"

// O(log n) kali, karena iterasi sebanyak n, dan n dibagi 2 setiap iterasinya
func pow(x, n int) int {
	xPowN := 1
	for n > 0 {
		// cek n ganjil atau genap
		isEven := n%2 == 0
		// jika ganjil x*x*x
		// jika genap x*x
		if !isEven {
			xPowN *= x
		}
		x *= x

		n /= 2
	}
	return xPowN
}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
