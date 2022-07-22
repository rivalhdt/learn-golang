package main

import "fmt"

func main() {
	var studentScore int = 80
	var huruf string
	if studentScore >= 80 {
		huruf = "A"
	} else if studentScore >= 65 {
		huruf = "B+"
	} else if studentScore >= 50 {
		huruf = "B"
	} else if studentScore >= 35 {
		huruf = "C"
	} else if studentScore >= 0 {
		huruf = "D"
	}
	fmt.Println(huruf)
}
