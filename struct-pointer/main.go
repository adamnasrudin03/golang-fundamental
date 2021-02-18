package main

import "fmt"

type Student struct{
	ID int
	Name string
	GPA float32
}

func (student *Student) Graduate3() {
	student.Name = student.Name + " S.T"
}

func main() {
	student := Student{1, "Adam Nasrudin", 3.45}
	fmt.Println(student.Name)
	Graduate(&student)
	fmt.Println(student.Name)
	
	student3 := Student{3, "Adam ajah", 3.45}
	fmt.Println(student3.Name)
	student3.Graduate3()
	fmt.Println(student3.Name)


	fmt.Println("===Beda alamat===")
	student2 := Student{2, "Adam", 3.45}
	fmt.Println("Alamat sebelum masuk func : ", &student2.Name)
	fmt.Println("Nilai sebelum masuk func : ",student2.Name)

	Graduate2(student2)

	fmt.Println("Alamat setelah masuk  func : ", &student2.Name)
	fmt.Println("Nilai setelah masuk  func : ",student2.Name)

}

func Graduate(student *Student) {
	student.Name = student.Name + " S.T"
}

func Graduate2(student Student) {
	student.Name = student.Name + " S.T"
	fmt.Println("Alamat didalam func: ", &student.Name)
	fmt.Println("Nilai didalam func: ",student.Name)
}