package main

import (
	"fmt"

	"exmaple.com/greet"
)

func main() {
	fmt.Println("Welcome to server")
	message := greet.Hello("sachin")

	fmt.Println(message)

}
