package main

import "fmt"

func main() {
	fmt.Println("Hello flow...")
	height := 20

	if height < 10 {
		fmt.Println("its less")
	} else if height == 20 {
		fmt.Println("Its too hign")
	} else {
		fmt.Println("Ist another")
	}

	// traditional way of if
	length := 20
	if length > 20 {
		fmt.Println("its too big")
	}

	// the new way for go
	// if statement;condition{}
	if length := 30; length > 29 {
		fmt.Println("Its greater than 29")
	}

}
