package main

import "fmt"

func main() {
	// while loop
	/*i:=1
	for i<3{
		fmt.Println(i)
		i=i+1

	}*/
	// infinite loop
	/*
		for {
			fmt.Println("arpit")
		}*/
	// classical for loop
	for i := 0; i <= 3; i++ {
		// break
		if i == 2 {
			continue
		}
		fmt.Println(i)

	}
	fmt.Println() // print a new line
	//range based for loop
	for i := range 5 {
		fmt.Println(i) // print from 0 to 4
	}
}
