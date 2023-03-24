package main

import "fmt"



func main (){
fmt.Println("Welcome to pointers")



number :=45

var intPointer *int 


intPointer=&number

*intPointer =*intPointer*2


fmt.Println(intPointer)
fmt.Println(*intPointer)
fmt.Println(number)





}