package main

import "fmt"




func main()  {

	fmt.Println("Maps in golang")


	languages:= make(map[string]string)

	languages["JS"] ="Javascript"
	languages["DT"] ="Dart"
	languages["SW"] ="Swift"
	languages["KT"] ="Kotlin"


delete(languages,"KT")


// looping through map


for key, value := range languages{

fmt.Printf("The key is %v, the value is %v \n",key,value)



}



	fmt.Println(languages)

	
}