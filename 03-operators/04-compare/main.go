package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 5
	c := 10

	fmt.Println(a == b)
	fmt.Println(b == c)

	fmt.Println(5 != 10) //true
	fmt.Println(5 != 5)  // false

	fmt.Println(10 > 5) // true
	fmt.Println(5 > 5)  // false

	fmt.Println(5 < 10) // true
	fmt.Println(5 < 5)  // false

	fmt.Println(5 >= 5)  // true
	fmt.Println(5 >= 10) // false

	fmt.Println(5 <= 5)  // true
	fmt.Println(10 <= 5) // false
}
