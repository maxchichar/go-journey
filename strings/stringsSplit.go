/*
🟢 Exercise 1 — Easy

Write:

func GetWords(s string) []string

Return all words in a sentence.

Example:

"Go is awesome" → ["Go", "is", "awesome"]
*/

// SOLUTION

/*
package main

import (
	"fmt"
	"strings"
)

func GetWords(s string) []string {
	word := strings.Split(s, " ")

	for i := range word {
		ch := word[i]
		ch = strings.ToLower(ch)
	}
	return word
}

func main()  {
	fmt.Println(GetWords("Go is awesome"))
}
*/

/*
🟡 Exercise 2 — Easy+

Write:

func CountWords(s string) int

Count how many words are in the string.
*/

// SOLUTION

/*
package main

import (
	"fmt"
	"strings"
)

func CountWords(s string) int{
	count := 0
	part := strings.Split(s, " ")

	for _, word := range part {
		if word != "" {
			count++
		}
	}
	return count
}

func main()  {
	fmt.Println(CountWords("anime is good"))
}
*/

/*
🟠 Exercise 3 — Medium

Write:

func UpperAllWords(s string) string

Convert every word to uppercase.

Example:

"go is fun" → "GO IS FUN"
*/

// SOLUTION

/*
package main

import (
	"fmt"
	"strings"
)

func UpperAllWords(s string) string{
	return strings.ToUpper(s)
}

func main()  {
	fmt.Println(UpperAllWords("go is fun"))
}
*/