package main

import "fmt"

func main() {

	//strings
	var nameOne string = "Maxwell"
	var nameTwo = "Chibueze"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "chimereze"
	nameThree = "charles"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameour := "sparkyechox"
	fmt.Println(nameour)

	//ints
	var ageOne int = 21
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne,ageTwo,ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint = 251

	fmt.Println(numOne,numTwo,numThree)

    var scareOne float32 = 25.98
	var scareTwo float64 = 12345678915372571523737373.7
    scoreThree := 1.7

	fmt.Println(scareOne,scareTwo,scoreThree)

}
