package main

import (
	"fmt"
	"io"
	// "io/ioutil"
	"os"
)

func main() {
	content := "some content"

	file, err := os.Create("./files.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println(length)
	defer file.Close()

	readFile("./files.txt")
}	

func readFile(filename string) {
	// dataInByte, err := ioutil.ReadFile(filename) depraceted now only calling os.ReadFile
	dataInByte, err := os.ReadFile(filename) 
	checkNilErr(err)
	fmt.Println("in bytes:", dataInByte)
	fmt.Println("in string: ", string(dataInByte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}