package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {

	// Struct with array

	employeeList := [3]employee{}

	employeeList[0] = employee{
		employeeID:   "101",
		employeeName: "John",
		phone:        "1234567890",
	}

	employeeList[1] = employee{
		employeeID:   "102",
		employeeName: "Doe",
		phone:        "1234567890",
	}

	employeeList[2] = employee{
		employeeID:   "103",
		employeeName: "Smith",
		phone:        "1234567890",
	}

	fmt.Println("employee1 = ", employeeList)

	// Struct with slice

	employeeSlice := []employee{}

	employee1 := employee{
		employeeID:   "101",
		employeeName: "John",
		phone:        "1234567890",
	}

	employee2 := employee{
		employeeID:   "102",
		employeeName: "Doe",
		phone:        "1234567890",
	}

	employeeSlice = append(employeeSlice, employee1)
	employeeSlice = append(employeeSlice, employee2)

	fmt.Println("employeeSlice = ", employeeSlice)

}
