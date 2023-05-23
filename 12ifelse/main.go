package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome")

	loginCount := rand.Intn(15)
	var result string

	if loginCount > 10 {
		result = "Regular User"
	} else if loginCount < 10 {
		result = "Not a regular user"
	} else {
		result = "Login Count = 10"
	}

	fmt.Println(loginCount, ":", result)
}
