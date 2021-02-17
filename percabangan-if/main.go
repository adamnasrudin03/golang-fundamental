package main

import "fmt"

func main() {

	// else if
	score := 110
	var grade string

	if score >= 85 && score <= 100 {
		grade = "A"
	} else if score >= 75 && score < 85 {
		grade = "B"
	} else if score >= 60 && score < 75 {
		grade = "C"
	} else if score >= 50 && score < 60 {
		grade = "D"
	} else if score  < 50 {
		grade = "E"
	} else {
		grade="Null"
	}

		fmt.Println("Grade : " , grade)
}