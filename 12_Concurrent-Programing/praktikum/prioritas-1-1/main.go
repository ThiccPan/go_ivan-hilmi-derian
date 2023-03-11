// Buatlah sebuah fungsi yang mencetak angka kelipatan x, dimana x merupakan parameter bilangan bulat positif. lalu jalankan fungsi tersebut dengan menerapkan goroutine, dengan interval waktu 3 detik!

package main

import (
	"fmt"
	"time"
)

func multiplied(x int, isDone chan bool) {
	fmt.Println("start multiplying", x)
	for i := 1; i <= 10; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(x, "*", i, "=", x*i)
	}
	isDone <- true
}

func main() {
	isDone := make(chan bool, 2)
	go multiplied(3, isDone)
	go multiplied(5, isDone)
	<-isDone
	<-isDone
}
