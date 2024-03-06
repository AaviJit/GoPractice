package main

import "fmt"

func main() {
	fmt.Println("Welcome to loop practice")
	var name []string
	name = append(name, "Avijit", "Mumu", "Shuva", "Okkhor", "Lopa")
	fmt.Println(name)

	for d := 0; d < len(name); d++ {
		fmt.Println(name[d])
	}

	for i := range name {
		if name[i] == "Mumu" {
			continue
		}
		fmt.Println(name[i])
	}

	for index, nam := range name {
		fmt.Printf("Index is %v and name is %v", index, nam)
		fmt.Printf("\n")
	}

	var count = 0

	for count <= 10 {
		if count == 7 {
			count++
			continue
		}
		fmt.Println(" count value is ", count)
		count++
	}

}
