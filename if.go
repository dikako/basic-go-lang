package main

import "fmt"

func main() {
	name := "andika"
	if name == "andika" {
		fmt.Println("Halo Fransiskus Andika Setiawan")
	} else if name == "fransiskus" {
		fmt.Println("Halo Fransiskus Andika Setiawan")
	} else {
		fmt.Println("Check nama anda")
	}

	//If using sort statement
	if length := len(name); length > 8 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama benar")
	}
}
