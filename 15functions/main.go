package main

import "fmt"

// cant write fxns inside another 
// can declare iife in golang too

func main() {
	fmt.Println("fxn")
	greet()

	func() {
		fmt.Println("iife")
	}()

	result := adder(1, 2)

	fmt.Println(result)

	pro := proAdder(1, 2, 3, 8)
	fmt.Println(pro)

	pro2, _ := proAdder2(1, 2, 3, 8) // return do kre to capture bhi krne pade
	fmt.Println(pro2)

}


func adder(val1 int, val2 int) int { // fxn signatures also needed to specify
	return val1 + val2
}

func proAdder(values ...int) int { // ...int are variadic fxns
	total := 0

	for _, value := range values {
		total += value
	}

	return total
}

func proAdder2(values ...int) (int, string) { // returning multiple types
	total := 0

	for _, value := range values {
		total += value
	}

	return total, "ho"
}

func greet() {
	fmt.Println("Namaste")
}

