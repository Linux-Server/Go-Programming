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

}
