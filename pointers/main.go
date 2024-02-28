package main

import "fmt"

func main() {
	fmt.Println("On pointers practice")
	var ptr *int
	fmt.Println("Value of ptr is ", ptr)

	var mynumber = 23
	var ptrr = &mynumber

	fmt.Println("Memory address ", ptrr)
	fmt.Println("Value in pointers  ", *ptrr)
	*ptrr = *ptrr * 3

	fmt.Println("Modified value is ", mynumber)

}
