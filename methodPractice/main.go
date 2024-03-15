package main

import "fmt"

func main() {

	user := User{"Avijit", "avijit_barua@hotmail.com", true, 32}
	fmt.Println("Printing entire user ", user)
	user.getEmail()
	user.setUserEmail("avijitbarua65@gmail.com")
	fmt.Println("Printing entire user after modification", user)
	user.getEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getEmail() {
	fmt.Println("Is user Active ", u.Email)
}

func (u User) setUserEmail(email string) {
	u.Email = email
}
