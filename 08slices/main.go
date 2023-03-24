package main

import (
	"fmt"
)





func main()  {
	fmt.Println("Welcome to slices")


var samplelist =[]string{"Apple","Mango","Peach","Jack Fruit"}

samplelist= append(samplelist, "Water melon","Grapes","Papaya")

// samplelist = append(samplelist[3:5])

// fmt.Printf("fruitlist is  %T\n",samplelist)
// fmt.Println(samplelist)
// fmt.Println(len(samplelist))


// var samplelist2 = make([]int,1)

// samplelist2[0]=99

// samplelist2 = append(samplelist2, 1,2,3,4,5,6,7,8,9,10)

// fmt.Println(samplelist2)
// fmt.Println(sort.IntsAreSorted(samplelist2))


//  sort.Ints(samplelist2)

// fmt.Println(samplelist2)


//delete peach or delete with an index


index:=3


samplelist = append(samplelist[:index],samplelist[index+1:]...)


fmt.Println(samplelist)



}