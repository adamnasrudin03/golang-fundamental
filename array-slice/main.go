package main

import "fmt"

func main() {
	//deklarasi array slice {Panjang array dinamis}
	var languages [] string
	//not working because  runtime error: index out of range [0] with length 0
	//languages[0] = "Go"

	//mengisi nilai array slice
	languages = append(languages, "go")
	languages = append(languages, "Java")

	//mengisi nilai array slice saat deklarasi
	gamesConsole := [] string{"Playstation 4", "Xbox One", "Nintendo Switch"}

	fmt.Println(languages)
	fmt.Println(gamesConsole)

	// jika index tidak dipakai maka harus diganti dengan _
	fmt.Println("For each	:")
	for _, console := range gamesConsole {
			fmt.Println(console)	
	}
	
}