package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to PIZAA MANIA!")
	fmt.Println("Please enter your rating...")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("You have given us, ", input)

	//conversion str to int
	rating, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	fmt.Println("Your rating is increased to ", rating+1)

	//conversion from str to float
	rating2, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	fmt.Println("In float value, ", rating2+1)

}
