package main

import (
	"fmt"
	"slices"
)

func main() {
	fruits1 := []string{"Mangga", "Apel", "Melon", "Pisang", "Sirsak", "Tomat", "Nanas", "Nangka", "Timun", "Mangga"}
	fruits2 := []string{"Bayam", "Wortel", "Kangkung", "Mangga", "Tomat", "Kembang Kol", "Nangka", "Timun"}

	fmt.Println(sliceFruits(fruits1, fruits2, 1))
}

func sliceFruits(fruits1 []string, fruits2 []string, operationType int)[]string{
	result := []string {}
	switch operationType{
		case 1 :	
			result = same(fruits1, fruits2)

		case 2 :
			result = difference(fruits1,fruits2)
	}

	return result
}

func same(fruits1 []string, fruits2 []string) []string{
	result := []string{}

	for _, v := range fruits1 {
		for _, w := range fruits2 {
			if v == w {
				if !slices.Contains(result, v) || len(result) == 0 {
					result = append(result, v)
				}
			}
		}
	}
	return result
}

func difference(fruits1 []string, fruits2 []string) []string{
	result := []string{}

	for _, v := range fruits1{
		for _, w := range fruits2{
			if v != w {
				if !slices.Contains(result, v) && !slices.Contains(fruits2, v) {
					result = append(result, v)
				}
			}
		}
	}

	return result
}