package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scanf("%d", &bilangan)
	for index := bilangan - 1; index > 0; index-- {
		if bilangan%index == 0 {
			fmt.Println(index)
		}
	}
}
