package main

import (
	"fmt"
	"math"
)

/*
*
Package math
- Package math merupakan package yang berisikan constant dan fungsi matematika
- https://golang.org/pkg/math/

math.Round(float64) -> Membulatkan float64 keatas atau kebawah sesuai dengan yang paling dekat
math.Floor(float64) -> Membulatkan float64 kebawah
math.Ceil(float64) -> Membulatkan float64 keatas
math.Max(float64, float64) -> Mengembalikan nilai float64 yang paling besar
math.Min(float64, float64) -> Mengembalikan nilai float64 yang paling kecil
*/
func main() {
	fmt.Println(math.Ceil(1.40))  // Result -> 2
	fmt.Println(math.Floor(1.60)) // Result -> 1
	fmt.Println(math.Round(1.60)) // Result -> 2
	fmt.Println(math.Max(1, 2))   // Result -> 2
	fmt.Println(math.Min(1, 2))   // Result -> 1
}
