package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to map practice")
	var mapVar = make(map[string]string)
	fmt.Println("Printing map after initializing ", mapVar)
	mapVar["JS"] = "JavaScript"
	mapVar["RB"] = "RUBY"
	mapVar["PY"] = "Python"
	mapVar["JA"] = "JAVA"

	fmt.Println("Printing map after initializing ", mapVar)
	fmt.Println("Js value from map =", mapVar["JS"])
	delete(mapVar, "PY")
	fmt.Println("After deleting ", mapVar)

	for key, value := range mapVar {
		fmt.Printf("Key value is %s %s \n", key, value)

	}

}
