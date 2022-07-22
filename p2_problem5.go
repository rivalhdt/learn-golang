package main

import "fmt"

func pangkat(base, pangkat int) int {
	var number int = 1
	for i := 1; i <= pangkat; i++ {
		number *= base

	}
	return number
}

func main() {
	fmt.Println(pangkat(2, 3))
	fmt.Println(pangkat(5, 3))
	fmt.Println(pangkat(10, 2))
	fmt.Println(pangkat(2, 3))
	fmt.Println(pangkat(7, 3))
}
