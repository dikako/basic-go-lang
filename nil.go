package main

import "fmt"

/**
- Biasanya did alam bahasa pemrograman lain, object yang belum diinisialisasi maka secara otomatis nilainya adalah null atau nil
- Berbeda dengan Go-Lang, di Go-Lang saat kita buat variable dengan tipe data tertentu, maka secara otomatis akan dibuatkan default valuenya
- Namu di Go=Lang ada data nil, yaotu data kosong
- Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map slice, pointer dan channel
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {
	person := NewMap("Dika")
	fmt.Println(person) // map[name:Dika]

	if person == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person["name"]) // Dika
	}

}
