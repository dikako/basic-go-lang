package main

import "fmt"

/**
Pass by value
- Secara default di Go-Lang semua variable itu di passing by value, buka by reference
- Artinya jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenrnya yang dikirim adalah duplikasi valuenya

Pointer
- Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
- Sederhananya, dengan kemampuan pointer kita bisa membuat pass by reference

Operator &
- Untuk membuat sebuah variable dengan nilai opinter ke variable yang lain, kita bisa menggunakan operator dan diikuti dengan nama variablenya

Operator *
- Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut
- Semua variable yang mengacu ke data yang sama tidak akan berubah
- Jika kita ingin mengubah seluruh variable yang mengacu kedata tersebut, kita bisa menggunakan operator *

Function new
- Sebelumnya untuk pointer dengan menggunakan operator &
- Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
- Namun function new hanya mengembalikan opinter ke data kososng, artinya tidak ada data awal
*/

type Address struct {
	kota, provinsi, negara string
}

func PassByValue() {
	address1 := Address{"Lampung Timur", "Lampung", "Indonesia"}
	address2 := address1

	address2.kota = "Metro"

	fmt.Println(address1) // {Lampung Timur Lampung Indonesia}
	fmt.Println(address2) // {Metro Lampung Indonesia}
}

func NotPassByValue() {
	address1 := Address{"Metro", "Lampung", "Indonesia"}
	address2 := &address1

	address2.kota = "Mesuji"

	fmt.Println(address1) // {Mesuji Lampung Indonesia}
	fmt.Println(address2) // &{Mesuji Lampung Indonesia}
}

func NotPassByValueReAssignVariableToNewValue() {
	address1 := Address{"Metro", "Lampung", "Indonesia"}
	address2 := &address1

	address2.kota = "Mesuji"

	address2 = &Address{"Sukadana", "Lampung", "Indonesia"}

	fmt.Println(address1) // {Mesuji Lampung Indonesia}
	fmt.Println(address2) // &{Sukadana Lampung Indonesia}
}

func NotPassByValueOverwirteAllVariableValue() {
	address1 := Address{"Metro", "Lampung", "Indonesia"}
	address2 := &address1

	address2.kota = "Mesuji"

	*address2 = Address{"Sukadana", "Lampung", "Indonesia"}

	fmt.Println(address1) // {Sukadana Lampung Indonesia}
	fmt.Println(address2) // &{Sukadana Lampung Indonesia}
}

func NewFunction() {
	address1 := Address{"Metro", "Lampung", "Indonesia"}
	address2 := new(Address)

	address2.kota = "Mesuji"

	fmt.Println(address1) // {Sukadana Lampung Indonesia}
	fmt.Println(address2) // &{Mesuji  }
}

func main() {
	PassByValue()
	NotPassByValue()
	NotPassByValueReAssignVariableToNewValue()
	NotPassByValueOverwirteAllVariableValue()
	NewFunction()
}
