package main

import (
	"fmt"
	"time"
)

/*
*
Package time
- Package time adalah package yang berisikan fungsionalitas untuk management waktu di GO-Lang
- https://golang.org/pkg/time

time.Now() -> Untuk mendapatkan waktu saat ini
time.Date(...) -> Untuk membuat waktu
time.Parse(layout, string) Untuk memparsing waktu dari string
*/
func main() {
	now := time.Now()
	fmt.Println(now.Local())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2023, time.April, 04, 25, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	parse, _ := time.Parse(time.RFC3339, "2023-01-02T15:04:05Z")
	fmt.Println(parse)
}
