package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

	// count := 10

	// var msg string

	// if count < 10 {

	// 	msg = "less than 10"
	// } else if count > 10 {
	// 	msg = "greater than 10"
	// } else {
	// 	msg = "exactly 10"
	// }
	// if 9%2 == 0 {
	// }
	// if count2 := 25; count2 < 10 {
	// }

	// fmt.Println(msg)

	//switch case

	// rand.Seed(time.Now().UnixNano())

	dicenumber := rand.Intn(6) + 1

	sampleslice := make([]int, 0)

	samplemap := make(map[int]string)

	switch dicenumber {
	case 1:
		fmt.Println("the number is : 1")
		fallthrough
	case 2:
		fmt.Println("the number is : 2")
		fallthrough
	case 3:
		fmt.Println("the number is : 3")
	case 4:
		fmt.Println("the number is : 4")
	case 5:
		fmt.Println("the number is : 5")
	case 6:
		fmt.Println("the number is : 6")
	}

	for i := 0; i < 10; i++ {

		sampleslice = append(sampleslice, i)
		samplemap[i] = "the value is " + strconv.Itoa(i)

	}

	for i, value := range samplemap {

		if i == 5 {
			goto sample
		}

		fmt.Println(i, value)
	}
sample:
	fmt.Println("this is a go to statement")
	fmt.Println(sampleslice)

}
