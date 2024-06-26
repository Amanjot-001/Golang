package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter")
	input, _ := reader.ReadString('\n')
	fmt.Println("readed", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1", numRating + 1)
	}
}
