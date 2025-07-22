package main

import "fmt"

func main() {
	// var value interface{} = "Hello, friend!"
	var value any = "Hello, all"
	fmt.Println(value)
	value = 50
	fmt.Println(value)

	if str, ok := value.(string); ok {
		fmt.Println("This is string", str)
	} else {
		fmt.Println("This is no string")
	}

	if i, ok := value.(int); ok {
		fmt.Println("This is int", i)
	} else {
		fmt.Println("This is no int")
	}

}
