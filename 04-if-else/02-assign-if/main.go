package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// num := rand.IntN(100)

	if num := rand.IntN(100); num > 50 { // вычисление в IF, num доступен исключительно в этом ветвлении
		fmt.Printf("Выпало число %d!\n", num)
	} else {
		fmt.Printf("Выпало число %d, маловато будет...\n", num)
	}

	// чаще всего используется для проверки ошибок
	//
	// if err:= doSomething(); err != nil {
	// 	fmt.Println("Произошла ошибка:", err)
	// }
	// ...

}
