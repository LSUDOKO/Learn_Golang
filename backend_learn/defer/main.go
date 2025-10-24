package main

import "fmt"

func print() {
	fmt.Println("counting numbers")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main() {
	defer fmt.Println("hello i am arpit")
	fmt.Println("hello ")
	print()
}
