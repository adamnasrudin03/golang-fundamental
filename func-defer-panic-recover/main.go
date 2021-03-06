package main

import "fmt"

func logging(){
	fmt.Println("Selesai memanggil function")
}
//Defer func adalah func yang bisa kita jadwalkan untuk dieksekusi setelah 
//sebuah func selesai di eksekusi.
//Defer func akan selalu dieksekusi walaupun terjadi error di func yang dieksekusi

func runApplication(value int){
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplication(10)
}
