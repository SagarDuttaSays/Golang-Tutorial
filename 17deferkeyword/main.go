package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer keyword in Golang")
	//defer keyword delays the statement to the last of the program
	//it will reverse the array/slice
	//it follows LIFO pattern
	values := make([]int, 5)
	values[0] = 1
	values[1] = 2
	values[2] = 3
	values[3] = 4
	values[4] = 5
	print(values...)

	defer fmt.Println("Heena")
	defer fmt.Println("Shraddha") //second last in, second out
	defer fmt.Println("Ketaki")   //last in, first out
}
func print(values ...int) {
	for _, val := range values {
		defer fmt.Print(val, " ")
	}
}
