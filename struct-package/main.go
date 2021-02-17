package main

import (
	"struct-package/management"
	"fmt"
)


func main() {
	user := management.User{ 3, "Adam", "Nasrudin", "Adam@gmail.com", true}
	user2 := management.User{ 3, "Alifah", "Nurdianti", "alifah@gmail.com", true}

	users := []management.User{user, user2}
	group := management.Group{"Group Belajar", user, users, true}
	
	fmt.Println(user.Display())
	fmt.Println(user2.Display())
	group.DisplayGroup()

}