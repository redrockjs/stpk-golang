package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 555

	s1 := strconv.Itoa(i)
	fmt.Println(s1)

	s2 := strconv.FormatInt(int64(i), 2)
	fmt.Println(s2)

	s3 := strconv.FormatInt(int64(i), 8)
	fmt.Println(s3)

	s4 := strconv.FormatInt(int64(i), 16)
	fmt.Println(s4)

	s5 := strconv.FormatInt(int64(i), 10) // strconv.Itoa(i)
	fmt.Println(s5)

	number := 10.5254

	fixed := strconv.FormatFloat(number, 'f', 2, 64)
	fmt.Println("Fixed number (f)", fixed)

	strn := fmt.Sprintf("Привет, вот число: %f", number)
	strc := fmt.Sprintf("%v", number)
	fmt.Println(strn, "\n", strc)

	sample := 62.231413
	tmp := strconv.FormatFloat(sample, 'f', 3, 64)
	fmt.Println(tmp)

}
