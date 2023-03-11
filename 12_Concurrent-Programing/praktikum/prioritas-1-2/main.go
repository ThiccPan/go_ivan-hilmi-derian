// Buatlah program yang mencetak bilangan kelipatan 3 dengan menerapkan goroutine dan channel!

package main

import (
	"fmt"
	"time"
)

func multiply3(n int, channel chan []int) {
	arr := make([]int, n)
	for i := 1; i <= n; i++ {
		time.Sleep(1 * time.Second)
		arr[i-1] = n * i
	}
	channel <- arr
}

func main() {
	multiplyBy := 5
	channel := make(chan []int)
	go multiply3(multiplyBy, channel)
	fmt.Println("kelipatan dari 3 sebanyak", multiplyBy, "adalah:", <-channel)
}