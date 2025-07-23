package main

import "fmt"

func main() {
	a := 5

	if a > 5 {
		fmt.Println("a > 5")
	} else if a == 5 {
		fmt.Println("a = 5")
	} else {
		fmt.Println("a не > 5")
	}

	//----------------------

	b, c := 7, 1

	if b > 5 && c < 3 {
		fmt.Println("b больше 5 и с меньше 3")
	}

	if b < 5 || c < 3 {
		fmt.Println("b меньше 5 или с меньше 3")
	}

	//----------------------

	d := 5

	if d > 0 {
		fmt.Println("d > 0")
		if d < 10 {
			fmt.Println("d < 10")
		}
	}

	//----------------------

	temp := -5

	if temp < 0 {
		fmt.Println("Город замерзает! Верните лето.")
	}

	if temp >= 0 && temp <= 35 {
		fmt.Println("Температура в норме. Продолжаем писать код.")
	}

	if temp > 35 {
		fmt.Println("Город в огне! Яичницу можно жарить на асфальте.")
	}

}
