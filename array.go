package main

import "fmt"

func main() {

	// Array start from 0 (ZERO)

	var country [5]string
	country[0] = "Indonesia"
	country[1] = "Singapura"
	country[2] = "Belanda"
	country[3] = "US"
	country[4] = "Kamboja"

	fmt.Println(country[0])
	fmt.Println(country[1])
	fmt.Println(country[2])
	fmt.Println(country[3])
	fmt.Println(country[4])

	var countryCode = [5]int{
		12345,
		23456,
		34567,
		45678,
		56789,
	}
	fmt.Println(countryCode)
	fmt.Println(countryCode[0])

	// Function Array
	// Operasi                  | Keterangan                           |
	// len(array)               | Untuk mendapatkan panjang array      |
	// array[index]             | Mendapat data di posisi index        |
	// array[index] = value     | Mengubah data di posisi index        |

	fmt.Println(len(country))
	fmt.Println(len(countryCode))

	//Test
	var A [3][5]string
	var B [5]string

	//for i := 1; i < 5; i++ {
	//	sum += i
	//}

	for i := 0; i <= 5; i++ {
		A = A[i]
	}
}
