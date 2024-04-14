package main

import "fmt"

func main() {
	day := []int{1, 2, 3, 4, 5, 6}

	for i := 0; i < len(day); i++ {
		fmt.Println(day[i])
	}

	for i := range day {
		fmt.Println(day[i])
	}

	for _, d := range day {
		fmt.Printf("%v\n", d)
	}

	x := 1
	for x < 10 {
		if x == 2 {
			goto lco
		}

		if x == 5 {
			x++
			continue
			// break
		}

		fmt.Println(x)
		x++
	}

	lco: 
		fmt.Println("jump")
}
