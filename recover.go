package main

import "fmt"

func endApps() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("Application selesai")
}

func runApps(error bool) {
	defer endApps()
	if error {
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApps(true)
}
