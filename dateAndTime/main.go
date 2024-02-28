package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is related to date time variable")
	timeVar := time.Now()
	fmt.Println(timeVar)
	fmt.Println("After formatting date is")
	fmt.Println(timeVar.Format("01-02-2006"))
	fmt.Println(timeVar.Format("15:04:05 Monday"))
	modifiedDate := time.Date(1971, time.December, 16, 11, 50, 0, 0, time.UTC)
	fmt.Println(modifiedDate.Format("01-02-2006 15:04:05 Monday"))
}
