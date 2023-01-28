package main

import "fmt"

func main() {
	/*
	* For
	* Perulangan di golang cuma ada for loops saja
	* Dalam bahasa pemrograman, biasanya da feature yang bernama perulangan
	* Salah satu feature perulangan adalah for loops
	 */

	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke: ", counter)
		counter++
	}

	/*
	* For dengan statement
	* Dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa ditambahkan for
	* Init statement, taiut statement sbeelum for di eksekusi
	* Post statement, yaitu yang akan selalu dieksekusi di akhir tiap perulangan
	 */

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	/*
	* For range
	* For bisa idgunakan untuk melakukan iterasi terhadap semua data collection
	* Data collection contohnya Array, Slice dan Map
	 */

	names := []string{"Satu", "Dua", "Tiga"}
	for index, name := range names {
		fmt.Println("Index ", index, "= ", name)
	}

	/*Println valuenya saja*/
	for _, name := range names {
		fmt.Println(name)
	}

	/* Cara akses sata Slice manual */
	slice := []string{"Satu", "Dua", "Tiga"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	/* Data maps */
	person := make(map[string]string)
	person["name"] = "Dika"
	person["title"] = "QA Engineer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
