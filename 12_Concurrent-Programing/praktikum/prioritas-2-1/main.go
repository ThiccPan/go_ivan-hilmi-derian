// Hitung frekuensi huruf dalam teks menggunakan perhitungan paralel (Bersamaan).

package main

import (
	"bufio"
	"fmt"
	"strings"
	"sync"
	"time"
)

func main() {
	data := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

	// map of unique char
	charList := make(map[string]int)

	// channel to receive data
	ch := make(chan map[byte]int, len(data)/5)

	// data reader
	sc := bufio.NewReaderSize(strings.NewReader(data), 5)
	var wg sync.WaitGroup

	// locker for read operation
	var locker sync.Mutex

	// concurrently split the str into arr and put val into channel
	for i := 0; i < len(data)/5-1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			b1 := make([]byte, 5)
			m1 := make(map[byte]int)

			// lock sc while reading
			locker.Lock()
			sc.Read(b1)
			locker.Unlock()
			
			for _, v := range b1 {
				m1[v]++
				fmt.Print(string(v))
			}

			ch <- m1
		}()
	}

	wg.Wait()
	close(ch)

	fmt.Println("channel val")
	for {
		charsArr, isPresent := <-ch
		if !isPresent {
			break
		}
		for idx, chars := range charsArr {
			charList[string(idx)] += chars
		}
	}
	fmt.Println(charList)
}
