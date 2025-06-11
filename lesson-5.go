package main

import "fmt"

func main() {
	// var ages [3]int = [3] int{20,25,30,35}
	var ages = [3]int{20, 25, 30}

	names := [4]string{"maxwell", "joy", "prisca", "favour"}
	names[1] = "maximaa"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (uses arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

}
