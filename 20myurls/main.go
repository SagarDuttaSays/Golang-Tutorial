package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghjs24032o42sf"

func main() {
	fmt.Println("Welcome to URLS in golang")
	fmt.Println(myurl)
	//extracting information from the URL
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of qparams: %T\n", qparams)
	// fmt.Println(qparams["coursename"])

	for key, val := range qparams {
		fmt.Printf("%v : %v\n", key, val)
	}
	//constructing URLs
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=Hitesh",
	}
	myurl2 := partsOfUrl.String()
	fmt.Println(myurl2)
}
