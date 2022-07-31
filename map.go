package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Fransiskus Andika Setiawan",
		"address": "Jakarta Pusat",
	}

	//Add data to map / ubah data map
	person["kost"] = "butini"

	fmt.Println(person)
	fmt.Println(person["address"])
	fmt.Println(person["name"])
	fmt.Println(person["kost"])
	fmt.Println(len(person))

	// Funtion map
	// Operasi									| Keterangan													|
	// len(map)									| Untuk mendapatkan jumlah data di map							|
	// map(key)									| Untuk data di map menggunakan key								|
	// map[key] = value							| Untuk mengubah data di map dengan key							|
	// make(map[TypeKey]TypeValue)				| Untuk membuat map baru										|
	// delete(map, key)							| Untuk menghapus data di map									|

	book := make(map[string]string)
	book["title"] = "Dika belajar golang"
	book["author"] = "Fransiskus Andika Setiawan"
	book["error"] = "Salah"
	fmt.Println(book)

	delete(book, "error")
	fmt.Println(book)
}
