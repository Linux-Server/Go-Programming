package main

import "fmt"

func main()  {
	fmt.Println("Hello variables")
	var username int 
	fmt.Println(username)
	fmt.Printf("The type of varibal is %T  \n", username)
    
	//implicit style
	var isLoggedIn  = true
	fmt.Println(isLoggedIn)

	// no var style
	number_of_users := 910
	fmt.Println(number_of_users)
}