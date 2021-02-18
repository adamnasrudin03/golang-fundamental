package main

import "fmt"

func main() {
	numberA := 5
	//numberB menyimpan alamat memori numberA [& = reference]
	numberB := &numberA

	fmt.Println(&numberA)
	fmt.Println(numberA)
	//tercetak alamat number A
	fmt.Println(numberB)
	//tercetak nilai dari alamat memori numberA [* = deference]
	fmt.Println(*numberB)

	*numberB = 10
	
	fmt.Println(numberB)
	fmt.Println(*numberB)
	fmt.Println(numberA)

	fmt.Println("=== Contoh 2 ===")
	var numberC int = 20
	var numberD *int = &numberC
	
	fmt.Println(&numberC)
	fmt.Println(numberC)
	fmt.Println(numberD)
	fmt.Println(*numberD)

	numberC = 30 
	
	fmt.Println(&numberC)
	fmt.Println(numberC)
	fmt.Println(numberD)
	fmt.Println(*numberD)


	
}