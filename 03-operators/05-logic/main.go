package main

import "fmt"

func main() {
	a := 5
	b := 10

	fmt.Println("a > 3 && b > 3 = ", a > 3 && b > 3) //true
	fmt.Println("a > 5 && b > 5 = ", a > 5 && b > 5) //false

	fmt.Println("a > 3 || b > 3 = ", a > 3 || b > 3)     //true
	fmt.Println("a > 5 || b > 5 = ", a > 5 || b > 5)     //true
	fmt.Println("a > 50 || b > 50 = ", a > 50 || b > 50) //false

	fmt.Println("!(a > 3)", !(a > 3)) //false
	fmt.Println("!(a > 5)", !(a > 5)) //true

	//---------------------------
	result1 := a > 3 || (b > 10 && a < 3)
	fmt.Println(result1) // true

	result2 := (a > 3 || b > 10) && a < 3
	fmt.Println(result2) // false

	//age := 30
	//role := "admin"
	//status := "paused"
	// age := 15
	// role := "user"
	// status := "active"
	age := 42
	role := "officer"
	status := "active"

	has18years := age >= 18
	isAdmin := role == "admin" || role == "moderator"
	isCorrectRole := (role == "admin" || role == "moderator" || role == "user")
	isActive := status == "active"

	hasAccess := isAdmin || (isCorrectRole && has18years && isActive)
	fmt.Println("hasAccess:", hasAccess)
}
