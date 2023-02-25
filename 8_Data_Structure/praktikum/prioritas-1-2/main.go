package main

import "fmt"

func Mapping(slice []string) map[string]int {
	mapped := make(map[string]int)
	for _, v := range slice {
		mapped[v]++
	}
	return mapped
}

func main() {
		fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
		fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
}