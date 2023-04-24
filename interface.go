package main

import "fmt"

/**
- Interface adalah tipe data abstract, dia tidak memiliki implementasi langsung
- Sebuah interface berisikan definisi-definisi method
- Biasanya interface digunakan sebagai kontrak

- Setiap tipe data yang sesua dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
- Sehingga kita tidak perlu mengimplementasikan interface secara manual
- Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kit aharus menyebutkan secara eksplist akan menggunakan interface mana
*/

type HashName interface {
	GetName() string
}

func SayHello(hashName HashName) {
	fmt.Println("Hello", hashName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	dika := Person{"Dika"}
	SayHello(dika)

	kucing := Animal{"Kucing"}
	SayHello(kucing)
}
