package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	week = append(week, "Sunday")

	for i := 0; i < len(week); i++ {
		fmt.Println("Today is", week[i])
	}

	seasons := []string{"Spring", "Summer", "Monsoon", "Autumn", "December"}

	for _, season := range seasons {
		fmt.Printf("This season is %v\n", season)
	}

	//break
	fmt.Println("break")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("i =", i)
	}
	//continue
	fmt.Println("continue")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("i =", i)
	}

	//goto
	fmt.Println("goto")
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto LCO
		}
		fmt.Println("i =", i)
	}
LCO:
	fmt.Println("You came to LCO section")
}
