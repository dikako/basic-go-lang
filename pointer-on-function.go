package main

import "fmt"

/*
*
Pointer di Function
- Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
- Oleh karena itu, jika mengubah data didalam function, data yang aslinya tidak akan pernah berubah
- Hal ini membuat variable asal menjadi tidak bisa dubah
- Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
- Untuk melakukan ini, kita juga bisa menggunakan pointer di function
- Untuk menjadikan sebuah parameter sebagai function, kita bisa menggunakan operator * di parameternya
*/
type AddressIndonesia struct {
	kota, provinsi, negara string
}

func ChangeAddressToIndonesia(address *AddressIndonesia) {
	address.negara = "Indonesia"
}
func main() {
	address := AddressIndonesia{"Sukadana", "Lampung", ""}
	fmt.Println(address) // {Sukadana Lampung }

	ChangeAddressToIndonesia(&address)
	fmt.Println(address) // {Sukadana Lampung Indonesia}
}
