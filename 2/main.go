package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(isAnagram("Ottwo", "Toto"))
}



func isAnagram(word1 string, word2 string) bool{

	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	Arrword1 := splitStringandSort(word1)
	Arrword2 := splitStringandSort(word2)


	if Equal(Arrword1, Arrword2) {
		return true
	}
	return false
}

func splitStringandSort(words string) []string{
	
	splitted := []string {}

	for _,v := range words{
		splitted = append(splitted, string(v))
	}
	sort.Strings(splitted)
	
	return splitted
}



func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}