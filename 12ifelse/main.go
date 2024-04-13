package main

import "fmt"

func main() {
	loginCnt := 23
	var result string

	if loginCnt < 10 {
		result = "Reg user"
	} else if loginCnt > 10 {
		result = "yo"
	} else {
		result = "ho"
	}

	fmt.Println(result)


	if num := 3; num < 3 {
		fmt.Println("hello")
	} else {
		fmt.Println("bye")
	}
}
