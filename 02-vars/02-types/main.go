package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*  Целые числа
	int
	int8
	int16
	int32
	int64

	uint
	uint8
	uint16
	uint32
	uint64

	float64

	complex64
	complex128

	string					// строка
	rune					// как char

	bool					// булев тип

	Составные типы
	Array [5]int			// массивы
	Slice					// динамические массивы
	Map                     // пары ключ-значение
	Struct					// структуры переменные типы данных

	Function				// для описания функций

	interface				// контракт где мы что-то передаем и у него будет определенный метод(функция) внутри

	pointer 				// указатель, указываем где в оперативной памяти лежат данные, указывает на область памяти где наши данные лежат
	*/

	v := 555

	fmt.Println(reflect.TypeOf(v))
}
