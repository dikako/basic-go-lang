package main

import "fmt"

// Solusi 1: Menggunakan looping
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

// Solusi 2: Menggunakan recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	//Solusi 1: Menggunakan looping
	fmt.Println(factorialLoop(10)) // Result: 3628800

	//Solusi 2: Menggunakan recursive function
	fmt.Println(factorialRecursive(10)) // Result: 3628800

	//Static
	fmt.Println(10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1) // Result: 3628800
}
