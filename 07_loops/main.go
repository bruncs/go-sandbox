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

	// FizzBuzz
	fmt.Println("\n\nFizzBuzz:")
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Printf("FizzBuzz ")
		} else if i%5 == 0 {
			fmt.Printf("Buzz ")
		} else if i%3 == 0 {
			fmt.Printf("Fizz ")
		} else {
			fmt.Printf("%d ", i)
		}
	}
}
