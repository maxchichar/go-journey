package main

import "fmt"

func main() {
	age := 21
	name := "Chibueze"

  //Print
  fmt.Print("hello, ")
  fmt.Print("World! \n")
  fmt.Print("new line \n")
  
  //Println
  fmt.Println("hello zentrillax!")
  fmt.Println("goodbye zentrillax!")
  fmt.Println("my age is", age, "and my names are", name)

  //Printf (formatted strings) %_ = to a format spercifier
  fmt.Printf("my age is %v and my name is %v \n", age, name)
  fmt.Printf("my age is %q and my name is %q \n", age, name)
  fmt.Printf("age is of type %T \n", age)
  fmt.Printf("you scored %f points! \n", 225.55)
  fmt.Printf("you scored %0.1f points! \n", 225.55)

  //Sprintf (save formatted strings)
  var info = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
  fmt.Println("the saved string is:", info)

}
