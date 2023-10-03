package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// fmt.Println("hello main")

	output, res := greetings.Hello("")

	fmt.Println(output, res)
}
