package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Dika koko", "Dika"))            // true
	fmt.Println(strings.Split("Dika koko", " "))                  // [Dika koko]
	fmt.Println(strings.ToLower("DIKA KOKO"))                     // dika koko
	fmt.Println(strings.ToUpper("dika koko"))                     // DIKA KOKO
	fmt.Println(strings.ToTitle("dika koko"))                     // DIKA KOKO
	fmt.Println(strings.Trim("dika koko      ", " "))             // dika koko
	fmt.Println(strings.ReplaceAll("dika koko", "dika", "dikas")) // dikas koko
}
