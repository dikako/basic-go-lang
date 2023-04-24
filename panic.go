package main

import "fmt"

func endApp() {
	fmt.Println("Application selesai")
}

func runApp(error bool) {
	defer endApp()
	// panic akan menghentikan program
	if error {
		panic("Aplikasi error")
	}
}

func main() {
	runApp(true)
}
