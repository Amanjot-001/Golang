package main

import "fmt"

const CapitalSoPublic string = "hleo" // capital letter so its public

// x := 5 not possible in global

func main() {
	var username string = "aman"
	fmt.Println(username)
	fmt.Printf("var is of type: %T \n", username)

	var x bool = false
	fmt.Println(x)
	fmt.Printf("var is of type: %T \n", x)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("var is of type: %T \n", smallVal)

	var smallFloat float32 = 255.55555548655
	fmt.Println(smallFloat)
	fmt.Printf("var is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVar int
	fmt.Println(anotherVar)

	// implicit type
	var str = "hello"
	fmt.Println(str)

	// no var style
	num:= 20
	fmt.Println(num)
}
