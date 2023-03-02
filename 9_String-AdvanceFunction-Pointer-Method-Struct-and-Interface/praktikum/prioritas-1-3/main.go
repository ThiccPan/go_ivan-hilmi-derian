package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}

	isContained := strings.Contains(a, b)

	if isContained {
		return b
	}

	return ""
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
