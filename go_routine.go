package main

import (
	"fmt"
	"time"
)

func slow(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s, ":", i)
	}
}
func main() {
	go slow("A")
	slow("B")
	time.Sleep(10 * time.Second)
	fmt.Println("all task done.")
}
