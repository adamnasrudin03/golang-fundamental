package main

import "fmt"

type  BangunDatar interface {
	HitungLuas() int
}

//Syarat / kontrak menggunakan interface harus menggunakan method di dalam interface
type Persegi struct {
	Sisi int
}
func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar int
}
func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {
	//Karena persegi dan persegi panjang menggunakan interface BangunDatar maka
	//dapat menggunakan fungsi LuasBangunDatar
	persegi := Persegi{Sisi: 5}
	luas := LuasBangunDatar(persegi)
	fmt.Println("Luas Persegi : ", luas)

	persegiPanjang := PersegiPanjang{Panjang: 5, Lebar: 10}
	luas = LuasBangunDatar(persegiPanjang)
	fmt.Println("Luas Persegi Panjang: ", luas)

	
}

func LuasBangunDatar(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}