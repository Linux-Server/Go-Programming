package main

import "fmt"

func main() {
	fmt.Println("Hello funcs")
	add()
	sub(10, 5)
	var mulRes = mul(10, 10)
	fmt.Println(mulRes)
	sugar(10, 10, "sachin")
	sam, ram := returnMultiple()      // retrun multiple syntax
	firstValue, _ := returnMultiple() // ignore the second value
	fmt.Println(sam, ram, firstValue)
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

//multiple return value should be wrapped with bracket

func returnMultiple() (string, int) {
	return "sachin", 100
}
