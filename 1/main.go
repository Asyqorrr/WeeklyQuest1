package main

import "fmt"

func main() {
	word := "bananas"
	removeDuplicateLetter(word)
}

func removeDuplicateLetter(words string) {
	
	unique := []string {}

	for i := range words{
		for i2 := range words {

			if words[i] != words[i2]{
				unique = append(unique, string(words[i]))

			}

		}
	}
	

	fmt.Println(unique)

}