package main

import "fmt"

var c, python, java bool

// outside we can not use i:=7 outside the main function
func main() {
	var i int
	k := 7
	fmt.Println(i, c, python, java, k)
}
