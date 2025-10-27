package main

import (
	"fmt"
	"time"
)

func loop(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	go loop("world")
	loop("hello")
}
