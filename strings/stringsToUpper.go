/*
📝 EXERCISE 1 — Easy

Write a function:

func ForceUpper(s string) string          <== use this function

It should return the uppercase version of the input.

Test it with:

"goLang"

"hello WORLD"

"123abc!"

*/

// SOLUTION

/*
package main

import (
	"fmt"
	"strings"
)

func ForceUpper(s string) string {
	return strings.ToUpper(s)
}

func main()  {
	fmt.Println(ForceUpper("goLang"))
	fmt.Println(ForceUpper("hello WORLD"))
	fmt.Println(ForceUpper("123abc!"))
}
*/

/*
📝 EXERCISE 2 — Medium
Write a function:

func IsAllUpper(s string) bool      <== use this function
Return true only if all letters are uppercase.

✔ "HELLO" → true
✔ "123!" → true (no lowercase letters)
✘ "Hello" → false
✘ "GO!" → true
✘ "gO" → false
*/

// SOLUTION

/*
package main

import (
	"fmt"
	"strings"
)

func IsAllUpper(s string) bool {
	if s == strings.ToUpper(s) {
		return true
	}else{
		return false
	}
}

func main()  {
	fmt.Println(IsAllUpper("HELLO"))
	fmt.Println(IsAllUpper("123!"))
	fmt.Println(IsAllUpper("Hello"))
	fmt.Println(IsAllUpper("GO!"))
	fmt.Println(IsAllUpper("gO"))
}
*/

/*
📝 EXERCISE 3 — Hard
Write a program that:

Reads a sentence.

Converts every word that is longer than 4 letters to uppercase.

Prints the transformed sentence.

Example:

Input:

this problem is really important
Output:

THIS problem is REALLY IMPORTANT
*/

// SOLUTION

package main

import(
	"fmt"
	"strings"
)

func upperTrans(sentence string) string{
	if len(sentence) <= 4 {
		return strings.ToUpper(sentence)
	}
	return sentence
}

func main()  {
	fmt.Println(upperTrans("this problem is really important"))
}