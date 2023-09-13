package main

import "fmt"

func main() {
	//long assignbment operator
	var smsSendingLimit string = "Sachin"
	//short assignment operatior
	username := "samray"

	//Example variables
	status := true
	amount := 13.99

	fmt.Printf("%v %v\n", smsSendingLimit, username)
	fmt.Println(status, amount)
	fmt.Printf("Hello %v\n", username)

	// print the type of the variables  - Use Printf
	fmt.Printf("The type of the vaaribale is %T\n", amount)

	// Initailize multiple variable in single line
	name, age, id := "Sachin", 22, 23.23
	fmt.Println(name, age, id)

}
