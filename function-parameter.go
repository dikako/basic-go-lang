package main

import "fmt"

func main() {
	firstName := "Fransiskus"
	lastName := "Andika"
	sayHelloTo("Dika", "Koko")
	sayHelloTo(firstName, lastName)
}

func sayHelloTo(firstName, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
