package main

import "fmt"

func main() {
	/*
		тип byte это алиас типа uint8 0..255
		var b byte = 200
		fmt.Println(b, string(b))

		s := "hi"

		fmt.Println(s[0], string(s[0]))
		fmt.Println(s[1], string(s[2]))
	*/

	//тип rune это алиас int32
	var r rune = '😁'
	fmt.Println(r, string(r))
}
