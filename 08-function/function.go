package main

import "fmt"

// Function ที่ไม่มีการรับค่าและส่งค่ากลับ
func hello() {
	fmt.Println("Hello foodborne.co.th")
}

func plus(value1 int, value2 int) int {
	return value1 + value2
}

func plus3value(value1, value2, value3 int) int {
	return value1 + value2 + value3
}

func main() {
	hello()
	result := plus(1, 2)
	fmt.Println(result)

	result2 := plus3value(5, 3, 4)
	fmt.Println(result2)
}
