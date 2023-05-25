package main

import "fmt"

//non parameterized function
func welcomeMessage() {
	fmt.Println("Namastey from INDIA")
}

func main() {
	fmt.Println("Welcome to functions in golang")

	//general function structure
	// func <func_name> (variable datatype) return_datatype {
	//body of the function
	// 	return <value>
	// }

	sum := adder(3, 5)
	fmt.Println("sum =", sum)
	welcomeMessage()

	var myslice = []int{1, 3, 5, 7, 9, 11, 13, 15}
	msg, prod := product(myslice...)
	fmt.Println(msg, prod)
}

//parameterized function
func adder(x int, y int) int {
	return x + y
}

func product(values ...int) (string, int) {
	p := 1
	for _, val := range values {
		p = p * val
	}

	return "The product is", p
}
