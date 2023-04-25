package main

import (
	"first/database"   // ini wajib memangil salah satu function yg ada dalam package
	_ "first/database" // ini jika kita ingin hanya menjalakn init function dari package lain atau tidak memngagil sama seklai function yg ada di package
	"fmt"
)

/*
*
Package Initialization
- Saat kita membuat package, kita bisa membuat sebuah function yang akan di akses ketika package kita diakses
- Ini sangat cocok ketika contohnya, jika paclage kita berisi function-function untuk berkomunikasi dengan database, kita membuat function inisialisasi untuk membuka koneksi ke database
- Untuk membuat function yang diakses secara otomatis ketika package itu di akses, cukup dengan membuat function dengan nama init

Blank Identifier
- Kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu function yang ada di package
- Secara default Go-Lang akan komplen ketika ada package yang diimport namun tidak digunakan
- Untuk menangani hal tersebut, kita bisa menggunakan blank identifier (_) sebelum nama package ketika melakukan import
*/
func main() {
	databaseConection := database.GetDatabase()
	fmt.Println(databaseConection)
}
