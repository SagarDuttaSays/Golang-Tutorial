package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your ID")
	input, _ := reader.ReadString('\n')
	fmt.Println("Your ID is", input)
}
