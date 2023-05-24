package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch case")
	rand.Seed(time.Now().Local().UnixMicro())
	die := rand.Intn(7)
	fmt.Println("die =", die)
	switch die {
	case 1:
		fmt.Println("Move 1 step")
	case 2:
		fmt.Println("Move 2 steps")
	case 3:
		fmt.Println("Move 3 steps")
	case 4:
		fmt.Println("Move 4 steps")
	case 5:
		fmt.Println("Move 5 steps")
	case 6:
		fmt.Println("Move 6 steps")
	default:
		fmt.Println("Invalid choice")
	}
	//fallthrough keyword
	var dayOfWeek int = 3

	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
		fallthrough
	case 2:
		fmt.Println("Tuesday")
		fallthrough
	case 3:
		fmt.Println("Wednesday")
		fallthrough
	case 4:
		fmt.Println("Thursday")
		fallthrough
	case 5:
		fmt.Println("Friday")
		fallthrough
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Day")
	}
}
