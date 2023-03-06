package main

import (
	"fmt"
	"strings"
)

type romanDigit struct {
	sym []string
	val []int
}

var romanList = romanDigit{
	sym: []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"},
	val: []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1},
}

func convToRoman(n int) {
	fmt.Println("input:", n)
	fmt.Print("output: ")
	for i, v := range romanList.val {
		if n/v > 0 {
			fmt.Print(strings.Repeat(romanList.sym[i], n/v))
			n -= v * (n/v)
		}
	}
	fmt.Println()

}

func main() {
	convToRoman(4)
	convToRoman(9)
	convToRoman(23)
	convToRoman(2021)
	convToRoman(1646)
	convToRoman(1646)

}

/*
Input: 4
Output: IV

Input: 9
Output: IX

Input: 23
Output: XXIII

Input: 2021
Output: MMXXI

Input: 1646
Output: MDCXLVI
*/
