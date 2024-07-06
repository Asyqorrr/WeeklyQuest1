package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int { 1, 2,7, 2, 3, 3, 4, 4, 4,5 }
	fmt.Println(removeDuplicate(nums))
}

func removeDuplicate(nums []int) []int{

	result := []int{}
	result = append(result, nums[0])

	for _, v := range nums {

		if !slices.Contains(result, v) {
			result = append(result, v)
		}
	}
	
	return result
}
