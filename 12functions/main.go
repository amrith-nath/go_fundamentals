package main

import "fmt"

func main() {
	samplefun()

	result := addnum(5, 6)

	fmt.Println(result)

	result2, message := addall(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)

	fmt.Println(result2, message)

}

func samplefun() {
	fmt.Println("This is a function sample")
}

func addnum(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

func addall(values ...int) (int, string) {

	sum := 0

	for _, value := range values {

		sum += value

	}

	return sum, "This is a simple message"
}
