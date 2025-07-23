package main

import "fmt"

func main() {
	day := 5

	switch day {
	case 1:
		fmt.Println("Понедельник")
	case 2:
		fmt.Println("Вторник")
	case 3:
		fmt.Println("Среда")
	case 4:
		fmt.Println("Четверг")
	case 5:
		fmt.Println("Пятница")
	case 6:
		fmt.Println("Суббота")
	case 7:
		fmt.Println("Воскресенье")
	default:
		fmt.Println("Неизвестный день")
	}

	switch {
	case day < 1:
		fmt.Println("Некоррекный день")
	case day >= 1 && day <= 5:
		fmt.Println("Рабочий день")
	case day >= 6 && day <= 7:
		fmt.Println("Выходной день")
	default:
		fmt.Println("Что у вас там происходит")
	}
}
