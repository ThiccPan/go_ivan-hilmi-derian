package main

import "fmt"

func simpleEquation(a, b, c int) {
	x, y, z := 0, 1, 2
	for ; x < 10000; x++ {
		for ; y < 10000; y++ {
			if y == z {
				continue
			}
			for ; z < 10000; z++ {
				if z == y {
					continue
				}
				if x+y+z == a && x*y*z == b && (x*x)+(y*y)+(z*z) == c {
					fmt.Println(x, y, z)
					return
				}
			}
			z = 2
			break
		}
	}
	fmt.Println("no solution")
}

func main() {
	simpleEquation(6, 6, 14)
}
