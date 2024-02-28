package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slice practice")
	var fruitList = []string{}
	fmt.Println("fruitlist is ", fruitList)
	fmt.Printf("fruitlist type %T\n ", fruitList)

	fruitList = append(fruitList, "Mango", "Banana", "Mealon")
	fmt.Println("Fruitlist is ", fruitList)
	fmt.Println(append(fruitList[1:2]))

	var sliceVar = make([]int, 3)
	fmt.Println("Initial slice var ", sliceVar)
	sliceVar[0] = 1
	sliceVar[1] = 2
	sliceVar[2] = 3

	fmt.Println("Printing slice var ", sliceVar)
	sliceVar = append(sliceVar, 4, 5, 6, 7, 8, 9)
	fmt.Println("Accomodating new var in slice ", sliceVar)

	var nameList = make([]string, 3)
	nameList[0] = "Shanu"
	nameList[1] = "Nihar"
	nameList[2] = "Dada"
	nameList = append(nameList, "Avijit", "Lopa", "Okkhor", "Shuva", "Mumu")
	fmt.Println("Namelist before sort ", nameList)
	sort.Strings(nameList)
	fmt.Println("Name list after sort", nameList)
	fmt.Println("Is list sorted =", sort.StringsAreSorted(nameList))

	// remove value from slice based on index
	nameList = append(nameList[:2], nameList[2+1:]...)
	fmt.Println(nameList)
	for val := range nameList {
		fmt.Println("In loop Name is ", nameList[val])
	}

}
