package main

import "fmt"

func main() {
	var x int = 64
	var y float64 = float64(x) // for x its give error
	//we cannot inplixitly change it type
	//we need to be expexitly type conversion
	fmt.Println(y)

}
