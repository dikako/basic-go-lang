package main

import "fmt"

func main() {
	name := "Fransiskus"
	if name == "Andika" {
		fmt.Println("Halo Andika")
	} else if name == "Fransiskus" {
		fmt.Println("Halo Fransiskus")
	} else {
		fmt.Println("Check nama anda")
	}

	//If using sort statement
	if length := len(name); length > 30 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama benar")
	}
}
