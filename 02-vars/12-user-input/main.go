package main

import (
	"fmt"
	"log"
)

func main() {
	//// var1
	// var firstName, lastName string
	// fmt.Print("Введите имя: ")
	// fmt.Scan(&firstName)
	// fmt.Print("Введите фамилию: ")
	// fmt.Scan(&lastName)

	//// var 2
	// fmt.Print("Введите имя и фамилию через пробел: ")
	// fmt.Scanln(&firstName, &lastName)
	// fmt.Printf("Вы ввели %s %s\n", firstName, lastName)

	fmt.Println("Введите текст:") // подготавливаем переменные для ввода
	var w1, w2, w3 string
	n, err := fmt.Scanln(&w1, &w2, &w3) // ожидаем ввода

	if err != nil { // проверяем, не было ли ошибки
		log.Fatal(err)
	}

	fmt.Printf("Количество считанных элементов: %d\n", n) // выводим данные
	fmt.Printf("Считанная строка: %s %s %s\n", w1, w2, w3)

	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Enter the string: ")
	// scanner.Scan()
	// str := scanner.Text()
	// fmt.Println("Your input: ", str)
}
