package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	diceNum := r.Intn(6) + 1
	fmt.Println(diceNum)
	// diceNum = 1

	switch diceNum {
	case 1:
		fmt.Println("hey 1")
		// fallthrough runs the next switch case also
	case 2:
		fmt.Println("hey 2")
		// fallthrough
	case 3:
		fmt.Println("hey 3")
	case 4:
		fmt.Println("hey 4")
	default:
		fmt.Println("no")
	}
}
