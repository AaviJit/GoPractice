package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to switch practice")
	var name = "Okkhor"
	testName(name)
	var number = rand.Intn(10)
	fmt.Println("Random number is ", number)
	testNumber(number)
}

func testName(name string) {
	switch name {
	case "Avijit":
		fallthrough
	case "Lopa":
		fmt.Println("This is for Avijit  and Lopa: ", name)
	case "Okkhor":
		fmt.Println("This is Okkhor: ", name)
	}
}

func testNumber(number int) {
	switch number {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fmt.Println("The number is :", number)
	case 4:
		fmt.Println("This is 4 : ", number)

	}

}
