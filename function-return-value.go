package main

import "fmt"

func main() {
	result := getHello("Dika")
	fmt.Println(result)
	fmt.Println(getHello(""))
}

func getHello(name string) string {
	if name == "" {
		return "Hello Bro name lu kosong"
	} else {
		return "Hello " + name
	}
}
