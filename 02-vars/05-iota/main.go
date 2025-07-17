package main

import "fmt"

const (
	Red = iota
	Green
	Blue
)

const (
	A = iota + 5
	B
	C
	D
)

const (
	_ = iota * 2
	K
	_
	M
)

const (
	_  = iota             // Игнорируем 0
	KB = 1 << (10 * iota) // 1 << (10 * 1) = 1024
	MB = 1 << (10 * iota) // 1 << (10 * 2) = 1048576
	GB = 1 << (10 * iota) // 1 << (10 * 3) = 1073741824
)

func main() {
	fmt.Println(Red, Green, Blue) // 0 1 2 похоже на enum в TypeScript
	fmt.Println(A, B, C, D)       // 5 6 7 8
	fmt.Println(K, M)             // 2 6

	fmt.Println(KB, MB, GB) // 1024 1048576 1073741824
}
