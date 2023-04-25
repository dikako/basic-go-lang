package main

import (
	"fmt"
	"strconv"
)

/**
Package strconv
- Sebelumnya kita sudah belajar cara konversi tipe data, misal dari int32 ke int34
- Bagaimana jika kita butuh melakukan konversi tipe datanya berbeda? Misal dari int ke stirng, atau sebaliknya
- Hal tersebut bisa kita lakukan dengan abntuan package strconv (string conversion)
- https://golang.org/pkg/strconv
*/

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean) // true
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number) // 1000000
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value) // 1000000

	valueInt, _ := strconv.Atoi("2000000")
	fmt.Println(valueInt) // 2000000
}
