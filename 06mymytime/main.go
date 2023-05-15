package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Current time: ", time.Now())
	fmt.Println("Current time in format: ", time.Now().Format("01-02-2006 Monday"))

	//create a date
	date := time.Date(2001, time.April, 24, 9, 24, 0, 0, time.Local)
	fmt.Println(date)
}
