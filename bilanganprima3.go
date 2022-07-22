package main

import "fmt"

func main() {
	var n int
	var b bool = false
	fmt.Scan(&n)
	fmt.Print(n)
	for i := 2; n > i; i++ {
		if n%i == 0 {
			b = true
		}
	}
}
