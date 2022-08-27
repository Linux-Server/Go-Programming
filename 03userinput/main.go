package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	var welcome string = "welcome to user Input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Plz enter the input value ?")
	//comma-ok  || comma-err syntax alternate for try cache

	input,_ := reader.ReadString('\n')
	fmt.Println("The input value is:", input)
	
}