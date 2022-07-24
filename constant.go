package main

import "fmt"

func main() {
	const firstName string = "Fransiskus"
	const lastName = "Setiawan"

	fmt.Println(firstName, lastName)

	const (
		class  = "IPA 3"
		number = 10
	)

	fmt.Println(class, number)
}
