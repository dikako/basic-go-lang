package main

import "fmt"

/**
Pointer di method
- Walaupun method akan menempel di struct, tapi sebenrnya data struct yang di akses di dalam method adalah pass by value
- Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method
*/

// ini contoh penggunaan yang buka pointer
type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	dika := Man{"Dika"}
	dika.Married()
	fmt.Println(dika.Name)
}
