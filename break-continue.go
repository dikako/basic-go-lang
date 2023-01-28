package main

import "fmt"

func main() {
	/*
	* Break & Continue
	* Break & Continue adalah kata kunci yang bisa digunakan dalam perulangan
	* Break digunakan untuk menghentikan seluruh perulangan
	* Continue adalah digunakan untuk menghentikan perulangan yang berjalan san langsung melanjutkan ke perulangan selanjutnya
	 */

	/* Break */

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Perulangan ke", i)
	}
	/* Result
	Perulangan ke 0
	Perulangan ke 1
	Perulangan ke 2
	Perulangan ke 3
	Perulangan ke 4
	*/

	/* Continue */
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}

		fmt.Println("Perulangan ke", i)
	}

	/* Result
	Perulangan ke 0
	Perulangan ke 1
	Perulangan ke 2
	Perulangan ke 3
	Perulangan ke 4
	***************Skip***************
	Perulangan ke 6
	Perulangan ke 7
	Perulangan ke 8
	Perulangan ke 9
	*/

}
