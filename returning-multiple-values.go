package main

import "fmt"

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	fmt.Println("lastName:", firstName)
	fmt.Println("firstName:", lastName)

	justFirstName, _ := getFullName()
	fmt.Println(justFirstName)
}

func getFullName() (string, string) {
	return "Dika", "Koko"
}
