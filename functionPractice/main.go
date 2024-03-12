package main

import "fmt"

func main() {
	sum1 := addNumber(11, 12)
	fmt.Println("Sum of two value is", sum1)

	sum2 := addMultipleNumbers(11, 12, 13, 14, 15, 13)
	fmt.Println("Sum of multiple value is", sum2)

	sum2, message := addMultipleNumbersWithMessage(11, 12, 13, 14, 15, 13)
	fmt.Println("Sum of multiple value is", sum2)
	fmt.Println("Returning message is", message)

}

func addNumber(value1 int, value2 int) int {
	return value1 + value2
}

func addMultipleNumbers(values ...int) int {
	var sum int = 0
	for i := 0; i < len(values); i++ {
		sum = sum + values[i]
	}
	return sum
}

func addMultipleNumbersWithMessage(values ...int) (int, string) {
	var sum int = 0
	for i := 0; i < len(values); i++ {
		sum = sum + values[i]
	}
	return sum, "You have added numbers successfully"
}
