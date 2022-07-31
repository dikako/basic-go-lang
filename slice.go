package main

import "fmt"

func main() {

	// Membuat Slice Dari Array
	// Membuat Slice        | Keterangan                                                             |
	// array[low:high]      | Membuat slice dari array dimulai index low sampai index sebelum high   |
	// array[low:]          | Membuat slide dari array dimulai index low sampai index akhir di array |
	// array[:high]         | Membuat slice dari array dimulai index 0 sampai index sebelum high     |
	// array[:]             | Membuat slice dari array dimulai index O sampai index akhir di array   |

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// Function Slice
	// Operasi                                   | Keterangan                                                                                                         |
	// len(slice)                                | Untuk medapatkan panjang                                                                                           |
	// cap(slice)                                | Untuk mendapatkan kapasitas                                                                                        |
	// append(slice, data)                       | Membuat slice baru dgn menambah data ke posisi terakhir slice, jika kapasitas penuh, makan akan membuat array baru |
	// make([]TypeData, length, capacity)        | Membuat slice baru                                                                                                 |
	// copy(destination, source                  | Menyalin slice dari source ke destination                                                                          |

	var slice1 = months[4:7]
	fmt.Println(slice1)      //[Mei Juni Juli]
	fmt.Println(len(slice1)) //3
	fmt.Println(cap(slice1)) //8

	months[5] = "Ubah Juni"
	fmt.Println(slice1)

	slice1[0] = "Ubah Mei"
	fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Bulan Dika")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Fransiskus"
	newSlice[1] = "Andika"
	fmt.Println(newSlice)      //[Fransiskus Andika]
	fmt.Println(len(newSlice)) //2
	fmt.Println(cap(newSlice)) //5

	// Pastikan kapastitas slicenya sama, supaya datanya tidak terpotong
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice) //[Fransiskus Andika]

	copySlice2 := make([]string, 1, cap(newSlice))
	copy(copySlice2, newSlice)
	fmt.Println(copySlice2) //[Fransiskus]

	//Hati hati dalam membuat array, perhatikan contoh di bawah ini
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniArray2 := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniArray2)
	fmt.Println(iniSlice)
}
