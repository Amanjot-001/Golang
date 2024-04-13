package main

import (
	"fmt"
	"sort"
)

func main() {
	var slice = []string{}
	fmt.Printf("type of slice is %T\n", slice)

	slice = append(slice, "mango", "banana")
	fmt.Println(slice)

	slice = append(slice[1:]) // first element is gone
	fmt.Println(slice)

	high := make([]int, 4)
	high[0] = 1
	high[1] = 2
	high[2] = 3
	high[3] = 4
	fmt.Println(high)

	// high[4] = 5 out of bound
	high = append(high, 6, 5) // memory re-accomadated
	fmt.Println(high)

	sort.Ints(high)
	fmt.Println(high)
	fmt.Println(sort.IntsAreSorted(high))
}
