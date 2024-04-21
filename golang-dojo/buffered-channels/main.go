package main

import "fmt"

func main() {
	channel := make(chan string)

	// go func() {
	// 	channel <- "first msg"
	// }()

	channel <- "first msg"

	fmt.Println(<-channel)
}