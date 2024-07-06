package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	nums1 := []int {1, 2, 3, 4, 5} 
	nums2 := []int {2, 4, 6, 7} 

	fmt.Println(sliceOperation(nums1,nums2,1))
}

func sliceOperation(nums1 []int, nums2 []int,typeOperation int)[]int{
	result := []int{}
	switch typeOperation{
		case 1:
			result = Difference1(nums1,nums2)
		case 2:
			result = Difference2(nums1,nums2)
		case 3:
			result = union(nums1,nums2)
		case 4:
			result = intersection(nums1,nums2)
	}
	return result
}

// A-B
func Difference1(nums1 []int, nums2 []int)[]int{
	result:= []int{}

	for _,v := range nums1{
		if !slices.Contains(nums2,v){
			result = append(result,v)

		}else{
			continue
		}
	}
	return result 

}

// B-A
func Difference2(nums1 []int, nums2 []int)[]int{
	result:= []int{}

	for _,v := range nums2 {
		if !slices.Contains(nums1, v){
			result = append(result,v)
		
		}else{
			continue
		}
	}
	return result

}

// AUB
func union(nums1 []int, nums2 []int) []int {
	result := []int {}

	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	// function duplicate nomor 2
	for _, v := range nums1 {
		if !slices.Contains(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// intersection
func intersection(nums1 []int, nums2 []int) []int {
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