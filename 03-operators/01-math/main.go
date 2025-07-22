package main

import (
	"fmt"
	"math"
)

func main() {
	random := 53.2 //math.Floor((rand.Float64()*100)*10) / 10

	tmp := math.Round(((random * 1.1) * 100) / 100)
	var isEven bool = false
	mod := int(tmp) % 2
	if mod == 0 {
		isEven = true
	}

	preLastDigit := math.Mod(random, 2)

	fmt.Printf("Исходное число: %v\n", random)
	fmt.Printf("Исходное число, увеличенное на 10%%: %v\n", tmp)
	fmt.Printf("Исходное число является четным: %v\n", isEven)
	fmt.Printf("Предпоследняя цифра целой части исходного числа: %v\n", preLastDigit)

}
