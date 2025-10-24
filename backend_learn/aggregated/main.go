package main

import "fmt"

// type vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	v := vertex{1, 2}
// 	// v.X = 21
// 	// fmt.Println(v.X)
// 	p := &v
// 	p.X = 1e5
// 	fmt.Println(v)
// }

type vertex struct {
	x, y int
}

var (
	v1 = vertex{1, 2}
	v2 = vertex{x: 21} //implicitly y=0
	v3 = vertex{}      //x=0 and y=0
	v4 = &vertex{1, 2} //has type *vertex
)

func main() {
	fmt.Println(v1, v2, v3, v4)
}
