package main

import(
	"fmt"
	"strings"
)

func is_palindrome(s string) bool{
	s = strings.ToLower(s)

	runes := []rune(s)

	left := 0
	right := len(runes) -1
	
	for left < right{
		if runes[right] != runes[left]{
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	word := "mama"
	fmt.Printf("Is %s a Palindrome ? %v\n", the, is_palindrome(word))
}