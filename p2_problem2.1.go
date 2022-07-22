package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)
	for index := 1; number >= index; index++ {
		if number%index == 0 {
			fmt.Println(index)
		}
	}
}
