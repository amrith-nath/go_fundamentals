package main

import "fmt"

func main() {
	defer fmt.Println("hello world1")
	defer fmt.Println("hello world2")
	defer fmt.Println("hello world3")
	defer fmt.Println("hello world4")

	sample()
	fmt.Println("simple hello")

}

func sample() {
	defer fmt.Println("hello world5")
	defer fmt.Println("hello world6")
	defer fmt.Println("hello world7")
	defer fmt.Println("hello world8")
}
