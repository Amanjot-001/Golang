package main

import "fmt"

func main()  {
	var arr [4]string
	arr[0] = "apple"
	arr[1] = "banana"
	arr[3] = "mango"

	// arr[2] ke liye 1 and 3 mai 2 space ka gap aagya
	// len is fixed 4 not matter ki kitne bhare h
	fmt.Println(arr, len(arr))

	var list = [3]string{"beans", "brocolli", "carrot"}
	fmt.Println(list)
}