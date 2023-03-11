// Buatlah program yang mencetak bilangan kelipatan 3 dengan menerapkan goroutine dan buffer channel!

package main

import (
	"fmt"
	"time"
)

func multiply3(x int, channel chan int) {
	defer close(channel)

	for i := 1; i <= x; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("3 *", i, "=", 3*i)
		channel <- i * 3
	}

	fmt.Println("data write done")
}

func main() {
	multiplyBy := 5
	channel := make(chan int, multiplyBy)

	go multiply3(multiplyBy, channel)
	numList := make([]int, 0)
	for {
		val, ok := <-channel
		// fmt.Println(val, ok)
		if !ok {
			break
		}
		numList = append(numList, val)
	}
	fmt.Println("kelipatan 3 sebanyak", multiplyBy, "kali adalah :", numList)
}
