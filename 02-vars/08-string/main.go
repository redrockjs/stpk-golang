package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "Hello"

	str1 := "Привет, мир"
	str2 := "Привет, \nмир"
	str3 := `Привет, 
	
	
	мир`

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)

	fmt.Println(len(str))
	fmt.Println("Runes count", utf8.RuneCountInString(str))

	strA := "   Hello!!!       "
	fmt.Println(len(strings.TrimSpace(strA)))
	upper := strings.ToUpper(strA)
	lower := strings.ToLower(strA)

	fmt.Println(upper, lower)
	fmt.Println(strings.Contains(strA, "llo"))

	message := "Go - это не просто язык, это СТИЛЬ ЖИЗНИ!"

	tmp := strings.TrimSpace(message)
	lowerm := strings.ToLower(tmp)
	fmt.Println(message, "\n", tmp, "\n", lowerm)
	fmt.Println(strings.HasPrefix(lowerm, "go"))
}
