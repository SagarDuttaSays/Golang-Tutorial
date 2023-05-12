package main

import "fmt"

func main() {
	var n int
	fmt.Println("Please enter the total number of elements in the array")
	fmt.Scanln(&n)
	// arr := make([]string, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Print("Please the ", i, "th element ")
	// 	fmt.Scan(&arr[i])
	// }
	// for i := 0; i < n; i++ {
	// 	fmt.Println("City = ", arr[i])
	// }
	city := [5]string{"Bhopal", "Vidisha", "Guna", "Ratlam", "Sagar"}
	for i, val := range city {
		fmt.Print(i, " ", val, " ")
	}
}
