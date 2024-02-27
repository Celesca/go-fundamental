package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("product = ", product)

	//add
	product["Macbook"] = 40000
	product["Ipad"] = 30000
	product["iPhone"] = 20000

	//delete
	delete(product, "Ipad")
	fmt.Println(product)

	//update
	product["Macbook"] = 45000

	//access
	value1 := product["Macbook"]
	fmt.Println(value1)

	courseName := map[string]string{"101": "Java", "102": "Python", "103": "C"}
	fmt.Println(courseName)
}
