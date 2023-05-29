package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("Welcome to Web Requests in Golang")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The response is of type:%T\n", response)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databytes))
}
