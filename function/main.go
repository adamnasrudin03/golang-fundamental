package main

import "fmt"

func main() {
	fmt.Println(add(10, 20))

	// luas, keliling := rectangle(10, 20)
	// fmt.Println("Luas Persegi panjang = " ,luas)
	// fmt.Println("Keliling Persegi panjang = " ,keliling)

	//tanda _ makaartinya tidak butuh/ tidak digunakan value ke 2
	luas, _ := rectangle(10, 20)
	fmt.Println("Luas Persegi panjang = " ,luas)
	
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