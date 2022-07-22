package main

import "fmt"

func main() {
	fmt.Print("Enter The Number: ")
	var bilangan int
	fmt.Scanf("%d", &bilangan)
	for index := 1; bilangan >= index; index++ {
		if bilangan%index == 0 {
			fmt.Println(index)
		}
	}
}
