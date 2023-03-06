package main

import "fmt"

func pascalCount(n int) [][]int {
	placeholder := make([][]int, 0)
	placeholder = append(placeholder, []int{1})
	for i := 1; i < n-1; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		prevCount := placeholder[i-1]
		for j := 1; j <= i/2; j++ {
			row[j] = prevCount[j] + prevCount[j-1]
			row[i-j] = prevCount[j] + prevCount[j-1]
		}
		placeholder = append(placeholder, row)
	}

	return placeholder
}

func main() {
	i1:=pascalCount(7)
	for _, v := range i1 {
		fmt.Println(v)
	}
	fmt.Println(7/2)
}

/*
step:
1. create placeholder arr
2. add first row with 1
3. iterate i time until i = input
	3.1. make new row
	3.2. fill first and last idx of row with 1
	3.3. iterate j time until j = i
	3.4. add 1 to first and last idx of row
	3.5. add prev row at idx j + idx (j-1) to curr row j
	3.6. add the row to placeholder
4. return placeholder
*/