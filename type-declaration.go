package main

import "fmt"

func main() {
	type noKtp string
	type merried bool

	var noKtpDika noKtp = "12234567890998989"
	var merriedStatus merried = true
	fmt.Println(noKtpDika, merriedStatus)
	fmt.Println(merriedStatus)
}
