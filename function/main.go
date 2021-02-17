package main

import (
	"errors"
	"fmt")

func main() {
	fmt.Println(add(10, 20))

	// luas, keliling := rectangle(10, 20)
	// fmt.Println("Luas Persegi panjang = " ,luas)
	// fmt.Println("Keliling Persegi panjang = " ,keliling)

	//tanda _ makaartinya tidak butuh/ tidak digunakan value ke 2
	// luas, _ := rectangle(10, 20)
	// fmt.Println("Luas rectangle = " ,luas)

	
	luas, kel := rectangle(10, 20)
	fmt.Println("return value : ")
	fmt.Println("Luas  = " , luas, "Keliling = " , kel)

	luas2, kel2 := rectangle2(10, 20)
	fmt.Println("Predefined return value : ")
	fmt.Println("Luas  = " , luas2, "Keliling = " , kel2)

	fmt.Println("=======")
	scores := []int{ 10, 5, 5, 10, 15 }
	fmt.Println(sum(scores))
	
	
	fmt.Println("=======")
	result, err := calculate(10, 5, "g")
	if err != nil {
		fmt.Println("Terjadi Kesalahan")
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
	
}

func add (number , numberTwo int) int {
	return number + numberTwo
}

//function multiple return
func rectangle(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	
	return luas, keliling
}

//Predefined return value
func rectangle2(panjang, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	
	return 
}

//func menjumlah didalam array
func sum(numbers [] int) int {
	var total int;
	for _, number := range numbers {
		total = total + number
	}
	return total
}

//func mengembalikan error
func calculate (number , numberTwo int, operation string ) (int, error){
	var result int
	var errorResult error
	switch operation {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	default :
		errorResult = errors.New("Unknown Operation")
	}
	return result, errorResult
}

//note: nama func di awalai kapital maka publik sedangakan huruf kecil maka private