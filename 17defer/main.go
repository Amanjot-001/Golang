package main

import "fmt"

// defer jese hi aae vo fxn ke last mai chali jati h
// lifo use hota yaha

func main() {
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello")
	myDefer()
}

func myDefer() {
	for i := 0; i<5; i++ {
		defer fmt.Println(i)
	}
}