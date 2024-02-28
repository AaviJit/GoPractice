package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input a number as rating with buf io NewReader")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error at reading line with input reader")
		return
	}
	var trimmedRating, _ = strconv.Atoi(strings.TrimSpace(input))
	trimmedRating += 1
	fmt.Printf("Input stream is %d \n", trimmedRating)

	fmt.Println("Input number as rating with fmt.scanner")

	var rating int
	_, err = fmt.Scan(&rating)
	fmt.Printf("After incrementing rating is %d", rating+1)

}
