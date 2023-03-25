package main

import "fmt"

func main() {
	total := sumAll(10, 10, 1, 2)
	fmt.Println(total)

	//Using slice
	numbers := []int{10, 10, 1, 2, 1}
	total = sumAll(numbers...)
	fmt.Println(total)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, numbers := range numbers {
		total += numbers
	}
	return total
}
