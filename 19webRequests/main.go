package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://httpbin.org/get"

func main() {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("res is %T\n", res) // *Http.Response , astrik means original not copy

	defer res.Body.Close() // it is callers responsibilty to close the connec

	dataBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataBytes))

}
