package main

import (
	"fmt"
	"time"
)

func process1(c chan string, data string) {
	c <- data // ส่งข้อมูลเข้าไปในช่อง

}

func main() {
	ch := make(chan string)  // ช่องในตัวส่งข้อมูล
	go process1(ch, "Data1") // บอก go ให้เป็น Routines
	fmt.Println(<-ch)        // รับข้อมูลจากช่อง (เป็นการรอข้อมูลจากช่อง
	time.Sleep(5 * time.Second)
}
