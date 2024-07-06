package main

import "fmt"

func main() {
	nums := []int{9, 9, 9, 9, 9, 9}

	fmt.Println(plusOne(nums))
}

func plusOne(nums []int) []int {
	lastIndex := len(nums) - 1

	// if nums[lastIndex] + 1 == 10 {
	// 	nums[lastIndex-1] = nums[lastIndex-1] + 1

	// }

	for i := lastIndex; i >= 0; i-- {

		if nums[i]+1 == 10 {

			if i == 0 && nums[i]+1 == 10 {
				nums = append([]int{1}, nums...)
				nums[i+1] = 0

			} else {
				nums[i] = 0
			}

		} else {
			nums[i]++
		}
	}

	return nums
}