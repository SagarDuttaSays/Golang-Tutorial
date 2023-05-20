package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Structs in Golang")
	std1 := Student{"Abhishek Maurya", 23, "4-A", 11, true, time.Now().Weekday()}
	fmt.Println("Student details:", std1)
	fmt.Printf("Student details: %+v", std1)
}

type Student struct {
	Name       string
	ID         int
	Class      string
	Age        int
	Attendance bool
	Day        time.Weekday
}
