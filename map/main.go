package main

import "fmt"

func main() {
	//deklarasi map[KeyType]ValueType
	var languageMap map[string]int
	languageMap = map[string]int{}

	languageMap["go"] = 9
	languageMap["java"] = 3
	languageMap["javascript"] = 10

	fmt.Println(languageMap)
	fmt.Println(languageMap["java"])

	//deklarasi & Langsung mengisi nilai map
	myMap := map[string] string{
		"go": "go language", 
		"java": "Kopi java",
		"ruby": "Ruby on Rails",
	}
	
	fmt.Println(myMap["java"])

	for key, value  := range myMap {
			fmt.Println("key : " , key , ", value : " , value)
	}

	fmt.Println("======")
	delete(myMap, "ruby")
	
	for key, value  := range myMap {
			fmt.Println("key : " , key , ", value : " , value)
	}

	fmt.Println("======")
	// Pengecekan value map
	//cara 1
	value := myMap["phyton"]
	fmt.Println(value)

	fmt.Println("cara 2")
	//cara 2
	value, isAvalaiable := myMap["phyton"]
	fmt.Println(isAvalaiable)
	fmt.Println(value)
}