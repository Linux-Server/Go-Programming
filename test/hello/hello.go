package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	fmt.Println("hello main")

	output := greetings.Hello("im happy here")

	fmt.Println(output)
}
