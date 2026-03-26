package main

import (
	"fmt"
	"time"
	"strings"
)

var word string
var userInput string

func indexProcessor(word string) string {
	fmt.Println("Welcome Indexio")
	fmt.Println()
	time.Sleep(3600 * time.Millisecond)

	fmt.Println("Loading...")
	time.Sleep(3 * time.Second)
	fmt.Println()

	fmt.Print("Input a word: ")
	fmt.Scanln(&word)
	fmt.Println()
	
	fmt.Println("Select an operation")

	operation1 := "1. lastword"
	operation2 := "2. capitalize"
	operation3 := "3. deleteIndex"

	fmt.Println(operation1)
	fmt.Println(operation2)
	fmt.Println(operation3)
	fmt.Println()

	fmt.Scan(&userInput)

	switch userInput {
	case "lastword":
		return string(word[len(word)-1])
	case "capitalize":
		return strings.ToUpper(word)
	case "deleteIndex":
		fmt.Println()
		fmt.Print("Input the index number you want to delete: ")

		var userDelIndex int
		fmt.Scanln(&userDelIndex)
		fmt.Println()

		LetWord := len(word)

		for {
			if userDelIndex < 0 || userDelIndex >= LetWord {
				fmt.Printf("Error input index within range %v: ", LetWord)
				fmt.Scanln(&userDelIndex)
				continue
			}
			break
		}

		emptyRune := []rune{}

		for i, ch := range word {
			if userDelIndex != i {
				emptyRune = append(emptyRune, ch)
			}
		}

		return string(emptyRune)
	default:
		return "Invalid Operation"
	}
}

func main()  {
	fmt.Println(indexProcessor(userInput))

}
