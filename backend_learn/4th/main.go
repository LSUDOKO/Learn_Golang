package main

import (
	"fmt"
	"math/cmplx"
)

var (
	value   bool       = false
	Maxint  uint64     = 25
	complex complex128 = cmplx.Sqrt(12 + 3i)
)

func main() {
	fmt.Printf("Type : %T value : %v\n", value, value)
	fmt.Printf("Type : %T value : %v\n", Maxint, Maxint)
	fmt.Printf("Type : %T value : %v\n", complex, complex)
}
