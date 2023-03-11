// Buatlah  program yang yang menerapkan mutex! Jenis program yang dibuat bebas (contoh: perhitungan faktorial).

package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeNumber struct {
	mu sync.Mutex
	v  int
}

func (sn *SafeNumber) set(v int) {
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.v = v
}

func (sn *SafeNumber) SafeFactorial() int {
	sn.mu.Lock()
	defer sn.mu.Unlock()
	val := sn.v
	for i:=1; i < sn.v; i++ {
		time.Sleep(1 * time.Second)
		val *= i
		fmt.Println(val)
	}
	return val
}

func main() {
	num := SafeNumber{}
	for i := 5; i <= 10; i++ {
		go num.set(i)
	}
	time.Sleep(5 * time.Second)
	num.SafeFactorial()
	
}
