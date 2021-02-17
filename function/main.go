package main

import "fmt"

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

//note: nama func di awalai kapital maka publik sedangakan huruf kecil maka private