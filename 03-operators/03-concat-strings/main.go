package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hello "
	str2 := "world"

	result := str1 + str2
	fmt.Println(result)

	result2 := fmt.Sprintf("%s%s", str1, str2)
	fmt.Println(result2)

	result3 := strings.Join([]string{str1, str2}, "")
	fmt.Println(result3)

	var buffer strings.Builder
	buffer.WriteString(str1)
	buffer.WriteString(str2)
	buffer.WriteString("!!!")

	result4 := buffer.String()
	fmt.Println(result4)
}
