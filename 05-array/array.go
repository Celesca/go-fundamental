package main

import "fmt"

var productName [4]string
var price [4]float64

func main() {
	productName[0] = "Macbook"
	productName[1] = "Ipad"
	productName[3] = "AirPods"
	productName[2] = "iPhone"
	fmt.Println(productName)

	price := [4]float32{4000, 3000, 2000, 200}
	fmt.Println(price)
}
