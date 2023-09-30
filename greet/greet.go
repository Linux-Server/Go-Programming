package greet

import "fmt"

func Hello(input string) string {
	fmt.Println("hello ", input)
	return "hello" + input
}
