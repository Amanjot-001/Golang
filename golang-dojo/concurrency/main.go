package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	defer func() {
		fmt.Println(time.Since(start))
	}()

	evilninjas := []string{"pam", "jim", "andy", "dwight"}

	for _, ninja := range evilninjas {
		go attack(ninja)
	}

	time.Sleep(time.Second * 2) // as the main is exiting before performing the goroutines
	// not a good way to handle, we dont know how much time these tasks take
}

func attack(target string) {
	fmt.Println("Throwing starts at", target)
	time.Sleep(time.Second) // finding the best way to attack
}