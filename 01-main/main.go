package main

import "fmt"

var numberInt, numberInt2 int = 1000, 2000

func main() {
	numberfloat := 25.4
	msg := "Hello"
	fmt.Println(numberInt)
	fmt.Println(numberInt2)
	fmt.Println(numberfloat)
	fmt.Println(msg)

	fmt.Println(float64(numberInt) + numberfloat)
	fmt.Println(msg + "World")
	fmt.Println("my money =", numberInt)
}
