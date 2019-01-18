package main

import "fmt"

func main() {
	// Long method
	i := 1
	for i <= 10 {
		fmt.Printf(" %d", i)
		i++
	}

	// Short method
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
}
