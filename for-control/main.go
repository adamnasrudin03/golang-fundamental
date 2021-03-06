package main

import "fmt"

func main() {

	for i:= 1; i <= 2; i++ {
		fmt.Println("Saya sedang belajar go : ", i)
	}
	fmt.Println("================================")

	//Break digunakan untuk menghentikan seluruh perulangan
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("[break] Perulangan ke ", i)
	}

	//Continue adalah digunakan untuk menghentikan perulangan yang berjalan,
	// dan langsung melanjutkan ke perulangan selanjutnya
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(" [continue]  Perulangan ke", i)
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