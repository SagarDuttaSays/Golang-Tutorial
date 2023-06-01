package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"course"` //name is aliased as course
	Price    int
	Plaform  string
	Password string   `json:"-"`              //hides the sensitive value
	Tags     []string `json:"tags,omitempty"` //omitempty will not show in the final output if the value is null
}

func main() {
	fmt.Println("Welcome to json in golang")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React JS", 199, "udemy.com", "12frg56", []string{"Front End", "JS"}},
		{"MERN stack", 399, "udacity.com", "r4324456", []string{"Full stack", "JS"}},
		{"React JS", 199, "codingninjas.com", "1234456", nil},
	}
	//pack this data as json data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"course": "React JS",
		"Price": 199,
		"Plaform": "udemy.com",
		"tags": [
				"Front End","JS"]
	}`)

	var lcoCourses course
	check := json.Valid(jsonDataFromWeb)
	if check {
		fmt.Println("JSON is VALID")
		//method 1
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
		//method 2
		//cases where you want to store data as key-value
		var mymaps map[string]interface{}
		json.Unmarshal(jsonDataFromWeb, &mymaps)
		fmt.Printf("%#v\n", mymaps)
		for key, val := range mymaps {
			fmt.Printf("%v : %v\n", key, val)
		}
	} else {
		fmt.Println("JSON not valid")
	}
}
