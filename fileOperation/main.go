package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var fileName = "myFile.txt"
	writeFile(fileName)
	readFile(fileName)
}
func writeFile(fileName string) {
	fmt.Println("Welcome to file practice...")
	var content = "Hi! This is Avijit Barua. " +
		"I am trying to create a new file. These are the content"
	file, err := os.Create(fileName)
	errCheck(err)
	length, err := io.WriteString(file, content)
	errCheck(err)
	fmt.Println("Length of file is ", length)
	err = file.Close()
	errCheck(err)
}

func readFile(fileName string) {
	bytes, err := os.ReadFile(fileName)
	errCheck(err)
	var content = string(bytes)
	fmt.Println(content)
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
