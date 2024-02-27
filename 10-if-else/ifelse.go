package main

import "fmt"

var score int

func main() {
	fmt.Println("grade calculator")
	fmt.Scanf("%d", &score)
	fmt.Println("score =", score)
	if score >= 80 {
		fmt.Println("A")
	} else if score >= 70 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else {
		fmt.Println("F")
	}
}
