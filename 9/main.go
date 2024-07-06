package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int {1, 2, 3, 4, 4,4,4,4, 4, 3, 3, 2, 4}
	fmt.Println(countFrequentElement(nums))
}

func countFrequentElement(nums []int)map[int]int{
	
	result := map[int]int{}

	unique := removeDuplicate(nums)

	for _,v := range unique{
		result[v] = 0
		for _,v2 := range nums{
			if v == v2 {
				result[v]++
			}
		}
	}

	return result

}

// function yg mirip seperti nomor 4
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