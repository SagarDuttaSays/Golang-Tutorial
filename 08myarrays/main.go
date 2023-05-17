package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Array session in Golang")
	var fruitList [5]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Cantaloupe"
	fruitList[4] = "Entawak"

	fmt.Println("Fruit List:", fruitList)

	var vegList = [5]string{"Potato", "Tomato", "Onion", "Aubergine", "Gourd"}
	fmt.Println("Vegetable List:", vegList)

	var spiceList [5]string
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 5; i++ {
		spiceList[i], _ = reader.ReadString('\n')
	}
	fmt.Println("Spice List is:", spiceList)
}
