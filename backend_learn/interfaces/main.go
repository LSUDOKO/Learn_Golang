// package main

// import (
// 	"fmt"
// 	"math"
// )

// type abser interface {
// 	Abs() float64
// }

// func main() {
// 	var as abser
// 	a := Myfloat(-math.Sqrt2)
// 	b := vertex{3, 4}
// 	as = a
// 	as = &b
// 	fmt.Println(as.Abs())

// }

// type Myfloat float64

// func (f Myfloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// type vertex struct {
// 	x, y float64
// }

// func (v *vertex) Abs() float64 {
// 	return math.Sqrt(v.x*v.x + v.y*v.y)
// }

package main

import "fmt"

type any interface{}

func main() {
	var i any = "hello world!"
	s := i.(string) //intialily i has type of interface now it has type of string
	fmt.Println(s)

	s, ok := i.(string) //here ok will check that is it true or false
	fmt.Println(s, ok)

	f, ok := i.(float64) //f will return 0 and ok will change false
	fmt.Println(f, ok)

}
