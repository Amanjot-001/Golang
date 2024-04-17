package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	defer func() {
		fmt.Println(time.Since(now))
	}()

	smokeSignal := make(chan bool)

	evilninja := "jim"
	go attack(evilninja, smokeSignal)

	// smokeSignal <- false // deadlock as same channel can't be used again

	fmt.Println(<-smokeSignal)
}

func attack(target string, attacked chan bool) {
	time.Sleep(time.Second)
	fmt.Println("throwing star at", target)
	attacked <- true
}