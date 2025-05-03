package main

import "fmt"
const roll=45
func main() {
	const name="arpit"
	// name="naman"// wrong
	// roll =25 // wrong
	fmt.Println(roll)
	fmt.Println(name)
	const (
		port =5000
		host ="localhost"
	)
	fmt.Println(port)
	fmt.Println(host)	
}