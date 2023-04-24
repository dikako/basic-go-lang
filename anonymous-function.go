package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	// ini yg di define dulu
	registerUser("Dika", blacklist)

	// ini juga bisa langsung dibuat dalam parameter valuenya
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
