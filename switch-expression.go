package main

import "fmt"

func main() {
	/*
	* Switch Expression
	* Selain if expression, untuk melakukan percabangan, kita juga bisa menggunakan Switch Expression
	* Switch expression sangat sederhana dibandingkan if
	* Biasanya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam satu varible
	 */

	name := "Fransiskus"

	switch name {
	case "Dika":
		fmt.Println("Hallo Dika")
	case "Fransiskus":
		fmt.Println("Hallo Fransiskus")
	case "Setiawan":
		fmt.Println("Hallo Setiawan")
	default:
		fmt.Println("Nama tidak terdaftar didalam switch case!")
	}

	/*
	* Switch dengan Short Statement
	* Sama dengan If, Switch juga mendukung short statement sebelum variable yang akan di check
	 */

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	/*
	* Switch Tampa Kondisi
	* Kondisi di switch expression tidak wajib
	* Jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut di setiap case nya
	 */

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu Panjang")
	case length > 5:
		fmt.Println("Nama lumayan Panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
