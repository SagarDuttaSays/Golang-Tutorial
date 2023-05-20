package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Maps in Golang")

	var states = make(map[string]string)

	states["UP"] = "Uttar Pradesh"
	states["TN"] = "Tamil Nadu"
	states["PB"] = "Punjab"
	states["WB"] = "West Bengal"
	states["GJ"] = "Gujarat"
	states["JH"] = "Jharkhand"
	states["AS"] = "Assam"
	states["JK"] = "Jammu & Kashmir"

	fmt.Println("Indian States")
	fmt.Println(states)

	delete(states, "JH")
	fmt.Println("Indian States")
	fmt.Println(states)

	_, found := states["JK"]
	fmt.Println("Jammu & Kashmir is part of India:", found)

	//looping
	for key, value := range states {
		fmt.Printf("%v : %v\n", key, value)
	}

}
