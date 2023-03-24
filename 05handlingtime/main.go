package main

import (
	"fmt"
	"time"
)





func main(){

fmt.Println("Time Study")

presentTime:= time.Now()

fmt.Println(presentTime)
fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))


createDate:=time.Date(2023,time.March,13,15,39,0,0,time.UTC)

fmt.Println(createDate.Format("01-02-2006, Monday"))


}