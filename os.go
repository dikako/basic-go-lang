package main

import (
	"fmt"
	"os"
)

/*
*
Package os
- Go-Lang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os
- Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan disemua sistem operasi)
- https://golang.org/pkg/os/
*/
func main() {
	args := os.Args
	fmt.Println(args) // go run os.go test -> [path test]
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname:", hostname)
	} else {
		fmt.Println("Error:", err)
	}

	username := os.Getenv("USER")
	fmt.Println(username)
}
