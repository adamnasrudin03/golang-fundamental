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



func endApp(){
	//Recover adalah function yang bisa kita gunakan untuk menangkap data panic
	//Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai")
}
//Panic func adalah func yang bisa kita gunakan untuk menghentikan program
//Panic func biasanya dipanggil ketika terjadi error pada saat program kita berjalan
//Saat panic func dipanggil, program akan terhenti, namun defer func tetap akan dieksekusi

func runApp(iniError bool){
	defer endApp()
	if iniError {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}


func main() {
	runApplication(10)
	runApp(true)
}
