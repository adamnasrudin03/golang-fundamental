package main

import "fmt"

type User struct{
	ID int
	FristName string
	LastName string
	Email string
	IsActive bool
	
}

type Group struct {
	Name string
	Admin User
	Users []User
	IsAvaliable bool
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

	users := []User{user, user2, user3}
	group := Group{"Group Belajar", user, users, true}


	display := displayUser(user)
	display2 := displayUser(user2)
	fmt.Println(display)
	fmt.Println(display2)
	displayGroup(group)

}

func displayGroup(group Group)  {
	fmt.Println("=====")
	fmt.Printf("Name : %s", group.Name )
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(group.Users))
	fmt.Println("")
	fmt.Println("Name Users : ")

	for _, user := range group.Users {
		fmt.Println(user.FristName)
	}

}

func displayUser(user User) string {
	//Sprintf u/ pengembalian nilai string
	result := fmt.Sprintf("Nama : %s %s, Email : %s ", user.FristName, user.LastName, user.Email)
	return result

}