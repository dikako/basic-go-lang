package main

import "fmt"

/**
- Go-Lang bukanlah bahasa pemmrograman yang berorientasi objek
- Biasanya dalam pemmrograman berorientasi object ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemmrograman tersebut
- Contoh di Java ada java.lang.Object
- Untuk menangani kasus seperti ini di Go-Lang kita bisa menggunakan inetrface kosong
- Interface kososng adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya

- Ada banyak contoh penggunaan interface kosong di Go-Lang, seperti:
-- fmt.Println(a ..interface{})
-- panic(v interface{})
-- recover() interface{}
-- dll
*/

func Ups() interface{} {
	return "Ups"
}
func main() {
	kosong := Ups()
	fmt.Println(kosong)
}
