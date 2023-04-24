package main

import "fmt"

func main() {
	counter := 0
	increment := func() {
		fmt.Println("increment")
		counter++
	}

	increment()          // increment 1
	increment()          // increment 2
	increment()          // increment 3
	increment()          // increment 4
	increment()          // increment 5
	increment()          // increment 6
	fmt.Println(counter) // Result: 6
}
