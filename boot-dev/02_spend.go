package main

import "fmt"

func main() {
	messageFromDoris := []string{
		"hello",
		"hai",
		"hoew",
	}

	numMessages := len(messageFromDoris)

	costPerMessage := 0.2

	total := costPerMessage * float64(numMessages)

	fmt.Printf("Dorsi spend %.2f dollar on monday", total)

}
