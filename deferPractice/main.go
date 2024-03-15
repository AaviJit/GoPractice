package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer practice")
	defer fmt.Println("Printing first number ", 1)
	defer fmt.Println("Printing second number ", 2)
	defer fmt.Println("Printing third number ", 3)
	defer fmt.Println("Printing forth number ", 4)

}
