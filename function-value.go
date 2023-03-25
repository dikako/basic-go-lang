package main

import "fmt"

func main() {
	sayGoodBye := getGoodBy
	fmt.Println(sayGoodBye("Dika"))
}

func getGoodBy(name string) string {
	return "Good Bye " + name
}
