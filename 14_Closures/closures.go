package main

import "fmt"

func clos() func() int {
	var cou int = 0
	return func() int {
		cou = cou + 1
		return cou
	}
}

func main() {
	imp := clos()
	fmt.Println(imp())
	fmt.Println(imp())

}
