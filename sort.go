package main

import (
	"fmt"
	"sort"
)

/*
*
Package sort
- Package sort adalah package yang berisikan utilitas untuk proses pengurutan data
- Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface
- https://golang.org/pkg/sort
*/

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Fransiskus", 24},
		{"Andika", 23},
		{"Setiawan", 22},
		{"Dika", 21},
		{"Koko", 20},
	}

	fmt.Println(users) // Result before sort -> [{Fransiskus 24} {Andika 23} {Setiawan 22} {Dika 21} {Koko 20}]

	sort.Sort(UserSlice(users))

	fmt.Println(users) // Result after sort -> [{Fransiskus 20} {Andika 21} {Setiawan 22} {Dika 23} {Koko 24}]
}
