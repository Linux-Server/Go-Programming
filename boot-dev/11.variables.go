package main

import "fmt"

func main() {
	var sendSmsLimit int
	var costPerSms float64
	var status bool

	fmt.Printf("%v %f %v ", sendSmsLimit, costPerSms, status)

	// how to format print the type

	fmt.Printf("%T", status)

	// declare multiple varibales on same line

	name, age := "Sachin", 28

	fmt.Println("the name and age is ", name, age)

	address, pincode := "Richmound road", 686743

	fmt.Printf("the address is %v and the pincode is %v", address, pincode)

	fmt.Println("")

	// coversion of types

	rollNumber := 11

	newRollNumber := float64(rollNumber)

	fmt.Printf("%f", newRollNumber)

}
