package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web request verbs in Golang")
	//GetRequest()
	//PostJSONRequest()
	PostformRequest()

}
func GetRequest() {
	const url = "http://localhost:1111/"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	//method 1
	var respString strings.Builder
	respString.Write(databytes)
	fmt.Println(respString.String())

	//method 2
	content := string(databytes)
	fmt.Println(content)
}

func PostJSONRequest() {
	const myurl string = "http://localhost:1111/post"
	//fake json payload
	requestBody := strings.NewReader(`
		{
		"name":"Sagar Dutta",
		"city":"Bhopal",
		"gender":"Male",
		"course":"BTech CSE with specialization in Health Informatics"
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	databyte, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(databyte))
	content := string(databyte)
	for idx, val := range content {
		fmt.Printf("%v : %v\n", idx, val)
	}
}

func PostformRequest() {
	const myurl string = "http://localhost:1111/post"

	//form data
	data := url.Values{}
	data.Add("firstname", "Sagar")
	data.Add("lastname", "Dutta")
	data.Add("college", "VIT Bhopal University")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	databyte, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(databyte))
}
