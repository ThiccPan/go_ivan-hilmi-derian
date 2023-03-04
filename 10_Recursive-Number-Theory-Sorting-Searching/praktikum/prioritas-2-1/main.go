package main

import "fmt"

// return all pair that match
func playingDomino(cards [][]int, deck []int) interface{} {

	arrPlayable := [][]int{}
	for i := 0; i < len(cards); i++ {
		if deck[0] == cards[i][0] || deck[1] == cards[i][0] {
			arrPlayable = append(arrPlayable, cards[i])
			continue
		}
		if deck[0] == cards[i][1] || deck[1] == cards[i][1] {
			arrPlayable = append(arrPlayable, cards[i])
			continue
		}
	}
	var playable interface{}
	if len(arrPlayable) == 0 {
		playable = "Tutup kartu"
		return playable
	}
	playable = arrPlayable
	return playable
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4,3}))
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3,6}))
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
}
