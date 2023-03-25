package main

import "fmt"

func main() {
	sayHelloWithFilter("Dika", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	sayHelloWithFilter2("Fransiskus", filter)
	sayHelloWithFilter2("Anjing", filter)
}

// Without Type declaration
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

// Without Type declaration

// Filter Type declaration for define parameter
type Filter func(string) string

func sayHelloWithFilter2(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

// With Type declaration

func spamFilter(name string) string {
	if name == "Anjing" {
		return "****"
	} else {
		return name
	}
}
