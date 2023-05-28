package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to file management in golang")

	content := "this is an a test document"
	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length is", length)
	defer file.Close()
}
