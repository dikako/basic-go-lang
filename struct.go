package main

import "fmt"

/**
- Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data
- Struct biasanya representasi data dalam program aplikasi yang kita buat
- Data di struct disimpan dalam field
- Sederhananya struct adalah kumpulan dari field
- Struct adalah template data atau prototype data
- Strict tidak bisa langsung digunakan
- Namun kita bisa membuat data/object dari struct yang telah kita buat
*/

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var dika Customer
	dika.Name = "Dika"
	dika.Address = "Lampung"
	dika.Age = 23

	fmt.Println(dika)         // {Dika Lampung, Indonesia 23}
	fmt.Println(dika.Name)    // Dika
	fmt.Println(dika.Address) // Lampung
	fmt.Println(dika.Age)     // 23

	// Struct literal
	dika2 := Customer{
		Name:    "Dika2",
		Address: "Lampung",
		Age:     23,
	}
	fmt.Println(dika2) // {Dika2 Lampung 23}

	// Atau bisa juga langsung define valuenya kayak gini
	dika3 := Customer{"Dika3", "Lampung", 23}
	fmt.Println(dika3) // {Dika3 Lampung 23}
}
