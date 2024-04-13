package main

import "fmt"

func main()  {
	var ptr *int
	ptr2 := new(int)
	fmt.Println(ptr, ptr2)

	num := 23

	var ptr3 = &num

	fmt.Println("actual memory add", ptr3)
	fmt.Println("value in that memory add", *ptr3)

	*ptr3 = *ptr3 + 2 // perfomed on actual var and not on copy
	fmt.Println(num)
}