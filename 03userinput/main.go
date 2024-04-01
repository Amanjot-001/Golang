package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "hello user"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter")

	// comma ok || error ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks", input)
	fmt.Printf("typs is %T", input)
}
