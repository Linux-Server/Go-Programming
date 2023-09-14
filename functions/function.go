package main

import "fmt"

func main() {
	fmt.Println("Hello funcs")
	add()
	sub(10, 5)
	var mulRes = mul(10, 10)
	fmt.Println(mulRes)
	sugar(10, 10, "sachin")
}

// create a func in go

func add() {
	fmt.Println("Im a add")
}

//pass parameter

func sub(x int, y int) {
	fmt.Println(x - y)
}

//return parameter

func mul(x int, y int) int {
	return x * y
}

// syntax sugar
func sugar(x, y int, name string) {
	fmt.Println(x, y, name)
}
