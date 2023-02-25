package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	arr2map := make(map[string]bool)
	arrMerged := make([]string, len(arrayA)+len(arrayB))
	uIdx := 0
	for _, v := range arrayA {
		if !arr2map[v] {
			arr2map[v] = true
			arrMerged[uIdx] = v
			uIdx++
		}
	}
	for _, v := range arrayB {
		if !arr2map[v] {
			arr2map[v] = true
			arrMerged[uIdx] = v
			uIdx++
		}
	}
	// bisa dioptimalkan space complexitynya jika return tidak harus berupa slice
	return arrMerged
}
func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
}
