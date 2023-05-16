package main

import "fmt"

func main() {
	//declaring a pointer
	var ptr *int
	fmt.Println("Declared a pointer", ptr)

	//initializing the pointer
	//*ptr will give the value stored at the address
	//ptr will give us the actual address
	//& is used for referrencing the address of a variable and store it into a pointer
	input := 54
	ptr = &input
	fmt.Println("Address of a pointer", ptr)
	fmt.Println("Value stored in a pointer", *ptr)
	fmt.Println("Address of a variable", &input)

	//value manipulation
	*ptr = *ptr + 4
	fmt.Println("Updated value check through pointer", *ptr)
	fmt.Println("Updated value check through variable", input)

}
