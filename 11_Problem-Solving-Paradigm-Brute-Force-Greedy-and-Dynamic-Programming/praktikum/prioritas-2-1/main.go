package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	energyArr := make([]int, len(jumps))
	for i := 1; i < len(jumps); i++ {
		cost1 := math.MaxInt
		cost2 := energyArr[i-1] + int(math.Abs(float64(jumps[i]-jumps[i-1])))

		if i > 1 {
			cost1 = energyArr[i-2] + int(math.Abs(float64(jumps[i])-float64(jumps[i-2])))
		}

		if cost1 < cost2 {
			energyArr[i] = cost1
		} else {
			energyArr[i] = cost2
		}
	}
	return energyArr[len(jumps)-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
