package main

import "fmt"

func main() {
	nums1 := []int{9, 2, 7}
	nums2 := []int{1, 3, 5}

	fmt.Println(addDigits(nums1, nums2))
}

func addDigits(nums1 []int, nums2 []int) []int {

	lastIndex := len(nums1) - 1

	// if nums[lastIndex] + 1 == 10 {
	// 	nums[lastIndex-1] = nums[lastIndex-1] + 1

	// }
	for i := lastIndex; i >= 0; i-- {
		if nums1[i]+nums2[i] >= 10 {

			if i == 0 && nums1[i]+nums2[i] >= 10 {
				nums1 = append([]int{1}, nums1...)
				nums1[i+1] = (nums1[i+1] + nums2[i]) % 10

			} else {
				nums1[i-1]++
				nums1[i] = (nums1[i] + nums2[i]) % 10
			}

		} else {
			nums1[i] = nums1[i] + nums2[i]
		}
	}
	return nums1
}