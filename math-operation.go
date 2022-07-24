package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	// Augmented Assignment
	// += (i = i + 10)
	// *= (i = i * 10)
	// -= (i = i - 10)
	// /+ (i = i / 10)
	// %= (i = i % 10)
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	// Unary Operator
	// ++ (i = i + 1)
	// -- (i = i - 1)
	// - (negative)
	// + (positive) but optional, because default number is positive
	// ! (Boolean kebalikan)
	i++ // i = i + 1
	fmt.Println(i)

	var (
		negative = -100
		positive = 100
	)

	fmt.Println(negative)
	fmt.Println(positive)
}
