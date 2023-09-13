package main

import "fmt"

var globalVariable string = "Nothing" //shot assignment operator dont work outside a function

func main() {
	//long assignbment operator
	var smsSendingLimit string = "Sachin"
	//short assignment operatior
	username := "samray"

	fmt.Println(globalVariable)

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

	//casting of type
	height := 178.7
	newCast := int(height) //cutting the fractions
	additionalCast := float64(newCast)
	fmt.Println(additionalCast)

	i := 42 // int
	g := 0.867 + 0.5i
	fmt.Printf("the compleNum type is %T %v\n", g, i)

}
