package main

import "fmt"

func sort(arr []int) {
	// split until arr len = 1
	if len(arr) <= 1 {
		return
	}
	m := (len(arr)) / 2
	l := make([]int, 0)
	l = append(l, arr[:m]...)
	r := make([]int, 0)
	r = append(r, arr[m:]...)
	sort(l)
	sort(r)
	merge(l, r, arr)
}

func merge(left, right, arr []int)  {
	lSize := len(arr)/2
	rSize := len(arr)-lSize
	i, l, r := 0, 0, 0
	for ;l < lSize && r < rSize;{
		if left[l] < right[r] {
			arr[i] = left[l]
			i++
			l++
		} else {
			arr[i] = right[r]
			i++
			r++
		}
	}
	for ;l < lSize; {
		arr[i] = left[l]
		i++
		l++
	}
	for ;r < rSize; {
		arr[i] = right[r]
		i++
		r++
	}
}

func MaximumBuyProduct(money int, productPrice []int) {
	// your code here
	sort(productPrice)
	for i, price := range productPrice {
		if money - price < 0 {
			fmt.Println("max item:",i)
			break
		}
		money -= price
	}
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000, 2000})      // 3
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}) // 4
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})   // 4
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})           // 1
	MaximumBuyProduct(0, []int{10000, 30000})                        // 0
}
