package otherpackage

/**
Access Modifier
- Di bahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier terhadap suatu function atau variable
- Di Go-Lang, untuk menentukan access modifier cukup dengan nama function atau variable
- Jika nama nya diawali dengan huruf besar, maka artinya bisa diakses dari pcakge lain, jika dimulai dengan huruf kecil maka artinya tidak bisa di access dari packgae lain
*/

var version = "1.0.0"      // tidak bisa diakses dari luar package
var Application = "golang" // bisa di akses dari luar package

// tidak bisa diakses dari luar package
func sayGoodBye(name string) string {
	return "Good bye " + name
}

// bisa di akses dari luar package
func SayHai(name string) string {
	return "Hai " + name
}
