package main

import "fmt"

func main() {

	for i:= 1; i <= 2; i++ {
		fmt.Println("Saya sedang belajar go : ", i)
	}
	fmt.Println("================================")

	//while di go tdk ada jdi pake for
	j := 1
	for j < 3 {
		fmt.Println("belajar go : ", j)
		j++
	}
	fmt.Println("================================")

	//For-each di go
	title := "Golang leanguange"
	for index, letter  := range title {
			fmt.Println("Index ke : " , index , ", letter : " , string(letter))
	}
	fmt.Println("================================")

	for index, buncis  := range title {
		if index %2 == 0 {
			fmt.Println("Index ke : " , index , ", letter : " , string(buncis))
		}
	}
	fmt.Println("================================")
	
	for index, letter  := range title {
		letterString :=  string(letter)

		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("Index ke : " , index , ", letter : " , string(letter))	
		}
	}
	
	
}