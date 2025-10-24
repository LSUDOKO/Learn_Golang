// package main

// import "fmt"

// func main() {
// 	names := [4]string{
// 		"arpit",
// 		"ayush",
// 		"dhruv",
// 		"somesh",
// 	}
// 	fmt.Println(names)
// 	var s []string = names[0:2] //
// 	var r []string = names[1:3]
// 	fmt.Println(s, r)
// 	r[0] = "XXX"
// 	fmt.Println(s, r)
// 	fmt.Println(names)

// }
package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10, 12, 14, 16}
	fmt.Println(s)
	t := []bool{true, false, true, true, false, true, false, true}
	fmt.Println(t)
	r := []struct {
		x int
		y bool
	}{
		{2, true},
		{3, false},
		{4, true},
		{5, false},
		{6, false},
		{7, true},
	}
	fmt.Println(r)

}
