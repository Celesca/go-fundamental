package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	data, _ := json.Marshal(&employee{101, "Sirasit", "0812345678", "sirasit@gmail.com"})
	fmt.Println(string(data))
}
