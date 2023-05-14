package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
)

func main() {
	fmt.Println("Welcome to math page")

	//using rand/math package
	// rand.Seed(time.Now().Unix())
	// r := rand.Intn(7)
	// fmt.Println("Random Number 1 is", r)

	//using rand/crypt package
	r2, _ := rand.Int(rand.Reader, big.NewInt(6))
	fmt.Println("Random Number 2 is", r2)
}
