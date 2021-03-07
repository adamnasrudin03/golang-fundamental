package main

import "fmt"

func main() {

	// deklarasi 1
	// var languages [5] string
	// languages[0] = "Go"
	// languages[1] = "Java"
	// languages[2] = "Javascript"
	// languages[3] = "Php"
	// languages[4] = "Phyton"

	// deklarasi 2
	// languages := [5] string{"Go", "Java", "Javascript", "Php", "Phyton"}

	//[...] berarti panjang array sesuai isinya
	languages := [...] string{"Go", "Java", "Javascript", "Php", "Phyton", "khotlin"}

	fmt.Println(languages)

	//length array
	fmt.Println(len(languages))

	for index, lang := range languages {
			fmt.Println("Index ke : " , index , ", Language : " , lang)	
	}


	
}