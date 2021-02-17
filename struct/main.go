package main

import "fmt"

type User struct{
	ID int
	FristName string
	LastName string
	Email string
	IsActive bool
	
}

func main() {
	user := User{}
	user.ID = 1
	user.FristName = "Adam"
	user.LastName = "Nasrudin"
	user.Email = "email@gmail.com"
	user.IsActive = true
	
	user2 := User{
		ID : 2,
		IsActive : true,
		FristName : "Alifah",
		LastName : "Nurdianti",
		Email : "alifah@gmail.com",
	}
	user3 := User{ 3, "Nama Depan", "Nama Belakang", "alifah@gmail.com", false}

	fmt.Println(user)
	fmt.Println(user2)
	fmt.Println(user3)

}