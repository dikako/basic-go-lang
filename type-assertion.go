package main

import "fmt"

/**
- Type Assertion merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
- Fitur ini sering kali digunakan ketika kita bertemu dengan data interface kosong

Type Assertion menggunakan switch
- Saat salah menggunakan type assertion, maka bisa berakibat  terjadi panic di aplikasi kita
- Jika panic dan tidak ter-recover, maka otomatis program kita akan berhenti
- Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertion
*/

func random() interface{} {
	return "OK"
}
func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	// dibawah ini adalah cara recover panic dari error
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}
