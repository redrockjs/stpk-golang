package main

import "fmt"

/* var (
	a = 5
	b = 2
	c int
) */

const (
	a = 5
	b = 2
	c = 100
)

// const DAYS_FOR_PAYMENT = 5 // for constants

func main() {
	// var age int = 30
	// var age = 30
	// age := 30
	// var x, y int = 1, 2
	// x, y := 1, 5
	// const Pi = 3.14 // PI, MY_VAR

	// daysForPayment := 5 // lower camel case
	// DaysForPayment := 5 // upper camel case for exports

	fmt.Println(a, b, c)
}
