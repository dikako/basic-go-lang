package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	// Defer function akan selelu di eksekusi walaupun function error
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println(result)
}

func main() {
	runApplication(0)
}
