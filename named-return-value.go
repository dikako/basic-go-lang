package main

import "fmt"

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Fransiskus"
	middleName = "Andika"
	lastName = "Setiawan"

	//return firstName, middleName, lastName => bisa didefinisikan secara expliciy tapi tidak wajib, bisa juga dengan hanya mengetikan return
	return
}
