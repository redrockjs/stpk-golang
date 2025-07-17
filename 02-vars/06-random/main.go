package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// randomNum := rand.Intn(100) // [0, 100]
	// fmt.Println(randomNum)

	// min := 10
	// max := 50
	// randomNum := rand.Intn(max-min) + min // [0, 100]
	// fmt.Println(randomNum)

	min := 1000
	max := 1000
	var random int = rand.IntN((max+1)-(min)) + min // [0, 100]
	fmt.Println(random)
}
