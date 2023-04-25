package main

import (
	"flag"
	"fmt"
)

/**
Package flag
- Package flag berisikan funsionalitas untuk memparsing command line argument
- https://golang.org/pkg/flag
*/

func main() {
	host := flag.String("host", "localhost", "Put your database host")            // default value localhost
	username := flag.String("username", "dika", "Put your database username")     // default value dika
	password := flag.String("password", "password", "Put your database password") // default value password

	// Re-assign value command -> go run flag.go -host=staging -username=staging -password=staginglohini -> result staging staging staginglohini

	flag.Parse()
	fmt.Println(*host, *username, *password) // localhost dika password
}
