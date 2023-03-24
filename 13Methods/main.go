package main

import "fmt"

func main() {

	amar := User{"amar", 26, true}

	fmt.Println(amar)

	amar.IsMajor()

}

type User struct {
	name string

	age int

	isOnline bool
}

func (u User) IsMajor() {
	if u.age < 18 {

		fmt.Printf("%v is not major yet", u.name)
	} else {
		fmt.Printf("%v is major", u.name)

	}
}
