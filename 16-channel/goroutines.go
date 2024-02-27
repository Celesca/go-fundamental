package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 1000; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	go f("Hello World!") // declare go routine
	go f("message2")     // declare go routine will concurrency will swap between two functions
	time.Sleep(5 * time.Second)
	// Go routine ทำให้ไม่ต้องรอ Process อันแรกที่นาน เราสามารถ Switch ตัวแรกขึ้นมาก่อนได้
}
