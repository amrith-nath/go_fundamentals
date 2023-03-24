package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main(){
fmt.Println("Enter a number here :")

reader := bufio.NewReader(os.Stdin)


input,err:=reader.ReadString('\n')



if err!=nil{

fmt.Println(err)


}else{

fnlInput,inputerr:= strconv.ParseFloat(strings.TrimSpace(input),64)


if inputerr!=nil{
	fmt.Println(inputerr)
}

	fmt.Println("here is your number",fnlInput +1)
}





}