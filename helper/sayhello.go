package helper

/**
- Dalam 1 package tidak boleh ada function yang sama

IMport
- Secara standar, file Go-Lang hanya bisa mengakses file Go-Lang lainnya yang berada dalam package yang sama
- Jika kita ingin mengakses file Go-Lang yang berada diluar package, maka kita bisa menggunakan import
*/
func SayHello(name string) string {
	return "Hello " + name
}
