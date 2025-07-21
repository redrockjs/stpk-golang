package main

import (
	"fmt"
	"math"
)

func main() {
	// random := rand.Float64()
	// fmt.Println(random)

	// min := 20.0
	// max := 50.0

	// random := rand.Float64()*(max-min+1) + min // +1 чтобы был включен максимум
	// fmt.Println(random)

	p := math.Floor(88.889*100) / 100

	fmt.Println(p)
}
