package main

import "fmt"

func main() {
	x := 42
	fmt.Println("Значение x:", x)
	fmt.Println("Указатель x:", &x)

	p := &x
	fmt.Println("Значение p:", p)
	fmt.Println("Значение p:", *p)

	*p = 100
	fmt.Println("Значение p:", *p)
	fmt.Println("Значение x:", x)

	var k *int
	fmt.Println(k)

	if k != nil {
		fmt.Println("Значение k:", *k)
	} else {
		fmt.Println("Указатель k равен nil:")
	}

	m := new(int)
	fmt.Println("Значение m:", m)
	fmt.Println("Указатель:", *m)
	*m = 50
	fmt.Println("Значение m:", *m)
}
