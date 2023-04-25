package main

import (
	"container/list"
	"fmt"
)

/*
*
Package container/list
- Package container/list adalah implementasi struktur data double linked list di Go-Lang
- https://golang.org/pkg/container/list/
*/
func main() {
	data := list.New()
	data.PushBack("Dika")
	data.PushBack("Fransiskus")
	data.PushBack("Setiawan")
	// Push Front -> untuk memasukan data ke paling awal
	// Push Back -> untuk memasukan data ke paling akhir

	fmt.Println(data) // &{{0x1400010e1b0 0x1400010e210 <nil> <nil>} 3}

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
		/**
		Dika
		Fransiskus
		Setiawan
		*/
	}

	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
		/**
		Setiawan
		Fransiskus
		Dika
		*/
	}
}
