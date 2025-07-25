package main

func main() {

	/*
		int8 - 1 байт  -127..127
		int16 - 2 байт -32768..32767
		int32 - 4 байт
		int64 - 8 байт


		uint - не может быть отрицательным

		uint8 - 1 байт  0..255
		uint16 - 2 байт 0..65535
		uint32 - 4 байт
		uint64 - 8 байт

		num := 4_324_654_654 	- вариант удобной записи разрядов
		num := 0b11001001 		- запись в двоичной системе исчисления
		num := 08056 			- запись в восьмеричной системе исчисления
		num := 0xA6F4 			- запись в шестнадцатиричной системе исчисления

		следует следить за переполнением, если данные переполняются то отсчет идет с низа типа данных, это сделано для производительности
	*/

	/*  Тип int в Go является удобным выбором, если не требуется экономия памяти.
	Он автоматически выбирает размер (32 или 64 бита) в зависимости от архитектуры
	системы. Если же необходимо оптимизировать использование памяти, следует выбирать
	более узкие типы, такие как int8, int16 или int32, в зависимости от диапазона значений.
	*/

}
