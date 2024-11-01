package main

import "fmt"

func findDuplicate(nums []int) int {
	var result int
	freqMap := make(map[int]int)
	for _, num := range nums {
		if freqMap[num] != 0 {
			result = num
			break
		}
		freqMap[num]++
	}
	return result
}

func main() {
	fmt.Println(findDuplicate([]int{1, 2, 3, 4, 4}))
	fmt.Println(findDuplicate([]int{1, 3, 3, 4, 5}))
	fmt.Println(findDuplicate([]int{1, 2, 3, 4, 6, 5, 5, 5}))
}
