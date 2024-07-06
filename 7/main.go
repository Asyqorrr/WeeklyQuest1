package main

import (
	"fmt"
	"slices"
)

func main() {
	nums1 := []int{1, 2, 7, 4, 7, 8}
	nums2 := []int{7, 7, 3, 2, 9}
	fmt.Println(sameElement(nums1, nums2))
}

func sameElement(nums1 []int, nums2 []int) []int{
	result := []int{}

	for _, v := range nums1 {
		for _, w := range nums2 {
			if v == w {

				if !slices.Contains(result, v) || len(result) == 0 {

					result = append(result, v)
				}

			}
		}
	}

	return result
}