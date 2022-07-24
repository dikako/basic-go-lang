package main

import "fmt"

func main() {

	// Boolean Operation
	// && (Dan)
	// || (Atau)
	//! (Kebalikan)

	// && Operation
	// Nilai 1 | Operator | Nilai 2 | Hasil |
	// true    | &&		  | true    | true  |
	// true    | &&		  | false   | false |
	// false   | &&		  | true    | false |
	// false   | &&		  | false   | false |

	// || Operation
	// Nilai 1 | Operator | Nilai 2 | Hasil |
	// true    | ||		  | true    | true  |
	// true    | ||		  | false   | true  |
	// false   | ||		  | true    | true  |
	// false   | ||  	  | false   | false |

	// ! Operation
	// Operator | Nilai 2 | Hasil |
	// !		| true    | false |
	// !		| false   | true  |

	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusAbsensi = absensi > 80
	fmt.Println(lulusNilaiAkhir)
	fmt.Println(lulusAbsensi)

	var lulus = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)

	// OR To the Point
	fmt.Println(nilaiAkhir > 80 && absensi > 80)
}
