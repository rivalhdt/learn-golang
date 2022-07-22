package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	for index := number - 1; index >= 1; index-- {
		if number%index == 0 {
			fmt.Println(index)
		}
	}
}
