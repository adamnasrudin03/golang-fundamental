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
	}
	
	fmt.Println(myMap["java"])
	
}