package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time in golang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // this has to be formatted with these values only

	// month is of type time.month here
	createdDate := time.Date(2020, time.September, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
