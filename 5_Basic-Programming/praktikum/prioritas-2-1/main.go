package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var nLoop int
	
	scanner.Scan()
	input := scanner.Text()
	fmt.Sscan(input, &nLoop)

	// starPureLoop(nLoop)
	starStringRepeat(nLoop)
}

// O(n^2) time complexity
func starPureLoop(nLoop int)  {
	// loop pertama mengatur ketinggian
	for i := 1; i <= nLoop; i++ {
		// loop kedua mengatur jumlah spasi sebelum bintang pertama diprint
		// jumlah spasi = input - current_iterasi_tinggi
		for j := 0; j < nLoop-i; j++ {
			fmt.Print(" ")
		}
		// loop ketiga mengatur jumlah bintang yang di print
		for k := 0; k < i; k++ {
			fmt.Print("* ")
		}
		// print line baru sebelum lanjut ke iterasi berikutnya
		fmt.Println()
	}
}

// O(n) time complexity
func starStringRepeat(nLoop int) {
	// loop mengatur ketinggian
	for i := 1; i <= nLoop; i++ {
        // print spasi sebelum bintang
		blank_space := strings.Repeat(" ", nLoop-i)
        fmt.Print(blank_space)
        // print bintang sebanyak i kali
		star_pattern := strings.Repeat("* ", i) 
        fmt.Println(star_pattern)
    }
}
