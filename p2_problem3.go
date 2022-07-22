package main

import "fmt"

func primeNumber(number int) bool {
	var result bool
	for i := 2; number > i; i++ {
		if number%i == 0 {
			result = false
			break
		} else {
			result = true
		}
	}
	return result
}

func main() {
	fmt.Println(primeNumber(11))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
