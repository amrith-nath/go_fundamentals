package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){

welcome:="Welcome to userInput"

fmt.Println(welcome)

reader := bufio.NewReader(os.Stdin)

fmt.Println("Enter a number")
input,err:=reader.ReadString('\n')

fmt.Println("the number is ",input)
fmt.Println("the error is ",err)



}