/*
🧪 EXERCISE 1 — Easy

Write:

func ForceLower(s string) string

Test with:

"HELLO"

"GoLang"

"123ABC!"

*/

//SOLUTION

/*
package main

import (
	"fmt"
	"strings"
)

func ForceLower(s string) string {
	return strings.ToLower(s)
}

func main() {
	fmt.Println(ForceLower("HELLO"))
	fmt.Println(ForceLower("GoLang"))
	fmt.Println(ForceLower("123ABC!"))
}
*/

/*
🧪 EXERCISE 2 — Medium

Write:

func IsAllLower(s string) bool

Rules:

Return true if all letters are lowercase

Examples:

"hello" → true
"123!" → true
"Hello" → false
"go!" → true
"Go" → false
*/

// SOLUTION

/*
package main

import (
	"fmt"
	"strings"
)

func IsAllLower(s string) bool {
	return s == strings.ToLower(s)
}

func main()  {
	fmt.Println(IsAllLower("hello"))
	fmt.Println(IsAllLower("123!"))
	fmt.Println(IsAllLower("Hello"))
	fmt.Println(IsAllLower("go!"))
	fmt.Println(IsAllLower("Go!"))
}
*/

/*
🧪 EXERCISE 3 — Hard

Write a function:

func NormalizeCommand(s string) string
Requirements:

Convert entire string to lowercase

BUT:

If a word is "error" → convert it to uppercase "ERROR"

Example:

Input:
"THIS is an ERROR and another Error"

Output:
"this is an ERROR and another ERROR"
*/

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func NormalizeCommand(s string) string{
	
// 	for i := 0; i < len(s); i++ {
// 		ch := s[i]
		
// 		if s == strings.ToLower(s) {
// 			if len(s) == 5 {
// 				return strings.ToUpper(string(ch))
// 			}
// 		}
// 	}
// 	return s
// }

// func main()  {
// 	fmt.Println(NormalizeCommand("THIS is an ERROR and another Error"))
// }