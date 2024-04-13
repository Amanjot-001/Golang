package main

import "fmt"

func main() {
	// no inheritance, no super, no parent in golang

	aman := User{"Aman", "mail@mail.com", true, 22}
	fmt.Println(aman)
	fmt.Printf("%+v\n", aman)

	fmt.Printf("Name is %v and Email is %v\n", aman.Name, aman.Email)
}

type User struct { // capital U in User and also inside caps first letter as they are public
	Name   string
	Email  string
	Status bool
	Age    int
}
