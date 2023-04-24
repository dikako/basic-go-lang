package main

import (
	"errors"
	"fmt"
)

/**
- Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interface nya adalah error

- Untuk membuat error, kita tidak perlu manual
- Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors (Package akan kita bahas secara detail di materi tersendiri)
*/

func Pembagaian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan nol tidak bisa bos!")
	} else {
		return nilai / pembagi, nil
	}
}
func main() {
	hasilBagi, err := Pembagaian(20, 0)
	if err == nil {
		fmt.Println("Hasil bagi:", hasilBagi)
	} else {
		fmt.Println("Error", err.Error())
	}

}
