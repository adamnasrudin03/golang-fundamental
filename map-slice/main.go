package main

import "fmt"

func main() {
	//Silce of map
	students := [] map[string] string{
		{
			"name": "Adam",
			"score": "A",
		},
		{
			"name": "Nasrudin",
			"score": "B",
		},
		{
			"name": "Agung",
			"score": "E",
		},
	}
	
	fmt.Println(students)
	fmt.Println("======")
	for _, student := range students {
		fmt.Println(student["name"] , " - " , student["score"])	
	}

}