package main

import "fmt"

func main() {
	// no inheritance, no super, no parent in golang

	aman := User{"Aman", "mail@mail.com", true, 22}
	fmt.Println(aman)
	fmt.Printf("%+v\n", aman)

	fmt.Printf("Name is %v and Email is %v\n", aman.Name, aman.Email)

	aman.GetStatus()
	aman.GetMail() // doesnt change original mail as a copy is passed

	fmt.Printf("Name is %v and Email is %v\n", aman.Name, aman.Email)

}

type User struct { // capital U in User and also inside caps first letter as they are public
	Name   string
	Email  string
	Status bool
	Age    int
	// oneAge int // not exportable now as not caps
}

func (u User) GetStatus() {
	fmt.Println(u.Status)
}

func (u User) GetMail() {
	u.Email = "newmail@g"
	fmt.Println(u.Email)
}
