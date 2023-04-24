package main

import "fmt"

/**
- Struct adalah tipe data seperti tipe data lainny, dia bisa digunakan sebagai parameter untuk function
- Namun jika kita ingin menambahklan method kedalam structs, sehingga seakan-akan sebuah struct itu memiliki sebuah function
- Method adalah function
*/

type CustomerPremium struct {
	Name, Address string
	Age           int
}

func (customer CustomerPremium) sayHello() {
	fmt.Println("Hello, My name", customer.Name)
}

func main() {
	dika := CustomerPremium{Name: "Dika"}
	dika.sayHello()
}
