package main

import (
	"fmt"
)

func distributeCandies(candyType []int) int {
	var res int
	maxCandies := len(candyType) / 2
	uniqueItems := make(map[int]int)
	for _, candy := range candyType {
		uniqueItems[candy] = 1
	}

	if uniqueItemsLen := len(uniqueItems); maxCandies > uniqueItemsLen {
		res = uniqueItemsLen
	} else {
		res = maxCandies
	}

	return res
}

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(distributeCandies([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
}
