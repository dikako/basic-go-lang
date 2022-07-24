package main

import "fmt"

func main() {

	// Comparison Operator
	// > (lebih dari)
	// < (kurang dari)
	// >= (lebih dari sama dengan)
	// <= (kurang dari sama dengan)
	// == (sama dengan)
	//!= (tidak sama dengan)

	var name1 = "Dika"
	var name2 = "Dika"
	var name3 = "Dikaa"
	var compareName1 = name1 == name2
	var compareName2 = name1 == name3
	fmt.Println(compareName1)
	fmt.Println(compareName2)

	var number1 = 20
	var number2 = 19
	var comparenumber = number2 < number1
	fmt.Println(comparenumber)
	fmt.Println(number1 > number2)
	fmt.Println(number1 == number2)
	fmt.Println(number1 != number2)
	fmt.Println(number1 >= number2)
	fmt.Println(number1 <= number2)

}
