package main

import (
	"first/helper"
	otherpackage "first/other-package"
	"fmt"
)

func main() {
	fmt.Println(helper.SayHello("Dika"))
	fmt.Println(otherpackage.SayHai("Other package"))
}
