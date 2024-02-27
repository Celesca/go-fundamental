package main

import "os"

func main() {
	data1 := []byte("Hello World\n")
	err := os.WriteFile("data.txt", data1, 0644)
	if err != nil {
		panic(err)
	}

	// Write File
	f, err := os.Create("employeeName.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	data2 := []byte("Sawit\n Manee")
	os.WriteFile("employeeName.txt", data2, 0644)
}
