package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}
func main() {
	fmt.Println("the sum of two number is ", add(4, 5))
}
