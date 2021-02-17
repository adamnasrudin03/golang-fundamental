package management

import "fmt"

type User struct{
	ID int
	FristName string
	LastName string
	Email string
	IsActive bool
	
}
//method User
func (user User) Display() string {
	return fmt.Sprintf("Nama : %s %s, Email : %s ", user.FristName, user.LastName, user.Email)

}

type Group struct {
	Name string
	Admin User
	Users []User
	IsAvaliable bool
}

//method Group
func (group Group) DisplayGroup()  {
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

// Penamaan struct, method, atau variabel diawali dengan :
// Kapital/uppercase berarti di export keluar package (public) / exported
// Huruf Kecil/lowercase berarti tidak di export keluar package (private) / unexported
