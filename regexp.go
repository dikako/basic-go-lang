package main

import (
	"fmt"
	"regexp"
)

/*
*
Package regexp
- Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
- Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
- https://github.com/google/re2/wiki/Syntax
- https://golang.org/pkg/regexp

regexp.MustCompile(string) -> Untuk membuat regex
Regexp.MatchString(string) bool -> Untuk mengecek apakah regex match dengan string
Regexp.FindAllString(string, max) -> Untuk mencari string yang match dengan maximum jumlah hasil
*/
func main() {
	regex := regexp.MustCompile(`[a-z]`)
	fmt.Println(regex.MatchString("dika"))                                      // Result -> true
	fmt.Println(regex.MatchString("diko"))                                      // Result -> true
	fmt.Println(regex.MatchString("DIKA"))                                      // Result ->  false
	fmt.Println(regex.FindAllString("dika 123 dikoko fasfas 35235 efwdfS", 10)) // Result -> [d i k a d i k o k o]

}
