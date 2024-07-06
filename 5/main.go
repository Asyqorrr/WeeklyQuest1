package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10} 
	fmt.Println(oddBeforeEven(nums))
}

func oddBeforeEven(nums []int) []int {

	odd := []int {}
	even := []int {}

	sort.Ints(nums)

	for _,v := range nums{
		if v%2 == 0{
			even = append(even, v)
		
		}else {
			odd = append(odd, v)
		}
	}

	odd = append(odd, even...)

	return odd
}