package main

import (
	"fmt"
)

func findErrorNums(nums []int) []int {
	var doubled, notPresent int

	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	for num, freq := range freqMap {
		if freq != 1 {
			doubled = num
			break
		}
	}

	for i := 1; i <= len(nums); i++ {
		if freqMap[i] == 0 {
			notPresent = i
			break
		}
	}

	return []int{doubled, notPresent}
}

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4, 5, 6}))
	fmt.Println(findErrorNums([]int{3, 2, 3, 4, 6, 5}))
	fmt.Println(findErrorNums([]int{1, 1}))
	fmt.Println(findErrorNums([]int{3, 3, 1}))
}
