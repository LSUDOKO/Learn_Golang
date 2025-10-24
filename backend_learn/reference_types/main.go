package main

import "fmt"

func main() {
	i, j := 42, 25
	p := &i //or var p *int
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 5
	fmt.Println(j)
}
