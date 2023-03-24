package main

import "fmt"







func main()  {
	user :=User{"amar","amar@g.com",26,true,30}

	user2 :=User{}

	user.Name="amarnath"
	user.Email="amar@g.com"
	user.Status=false
	user.Age=23
	user.minAge=3
	
	user2.Name="sample"
	user2.Email="sample@g.com"
	user2.Status=false
	user2.Age=98


	fmt.Printf("Details are : %+v",user)
	fmt.Printf("Details are : %+v",user2)

}




type User struct{

Name string 
Email string
Age int
Status bool

minAge int




}