package main

import "fmt"

func main() {
	fmt.Println("Welcome to the practice of stuck in go")
	var userVar = User{"Avijit", "avijit_barua@hotmail.com", true, 32}
	fmt.Println(userVar)
	fmt.Println(userVar.Name)
	fmt.Printf("Details of user is %+v \n", userVar)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
