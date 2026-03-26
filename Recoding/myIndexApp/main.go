//last word
//capitalise
//delete ask the user the index to delete and delete it. if the user put five delete the word at index 5, if index is 10 and the user input 25 it should tell the user that the index is out of range try again by restarting the program

package main

import (
	"fmt"
	"time"
	"strings"
)

var result string
var word string

func indexApp(word string) string {	
	fmt.Println("Welcome To My Index App")
	fmt.Println()
	time.Sleep(1 * time.Second)
	
	fmt.Println("Loading...") 
	time.Sleep(3 * time.Second)
	fmt.Println()
	fmt.Println()
	fmt.Println("Input a Word: ")
	fmt.Scanln(&word)
	fmt.Println()

	fmt.Println("Select An Operation")
	fmt.Println()

	operation1 := "Lastword"
	operation2 := "Capitalise"
	operation3 := "Delete-index"

	fmt.Println("1. ", operation1)
	fmt.Println("2. ", operation2)
	fmt.Println("3. ", operation3)
	fmt.Println()
	fmt.Scanln(&result)

	switch result {
	case "Lastword":
		return string(word[len(word)-1])
	case "Capitalise":
		return strings.ToUpper(word)
	case "Delete-index":
		fmt.Println()
		fmt.Println("Input index you want to delete: ")
		var deleteUserIndex int
		fmt.Scanln(&deleteUserIndex)
		fmt.Println()

		lWord := len(word)

		for {
			if deleteUserIndex >= lWord || deleteUserIndex < 0 {
				fmt.Printf("Index out of range!, input the an index within the range of %v: ", lWord)
				fmt.Scanln(&deleteUserIndex)
				continue
			}
			break
		}

		emptyRune := []rune{}

		for i, ch := range word {
			if deleteUserIndex != i {
				emptyRune = append(emptyRune, ch)
			}
		}

		return string(emptyRune)
	}
	return "Invalid Operation"
}

func main() {
	fmt.Println(indexApp(result))
}