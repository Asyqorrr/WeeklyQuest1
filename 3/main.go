package main

import (
	"fmt"
	"strings"
)

func main() {

	words := []string {"this","is","a","kanoha"}
	except := [] string {"is","a"}
	fmt.Println(capitalize(words, except))


}

func capitalize(words []string, except []string) []string{

	result := make([]string, len(words))

	for i, a := range words {
		// Check if the current element in A is in B
		found := false
		for _, b := range except {
			if a == b {
				found = true
				break
			}
		}

		// Capitalize the element if found in B
		if found {
			result[i] = a
		} else {
			result[i] = cap(a)
		}
	}
	return result
}

// return string dengan huruf pertama kapital
func cap(words string)string {
	return strings.ToUpper(words[0:1]) + words[1:]
}