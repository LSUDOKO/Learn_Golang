package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }
// func add(a, b int) int {
// 	return a + b
// }
// func getlanguages() (string, string, string, bool) {
// 	return "golang", "python", "java", true
// }
func processit() func(a int) int {
	return func(a int) int {
		return 4
	}
}
func main() {
	// lag1, lag2, lag3,lag4 := getlanguages()
	// fmt.Println(lag1, lag2, lag3,lag4)
	// fmt.Println(getlanguages())
	// result := add(12, 13)
	// fmt.Println(result)
	fn := processit()
	fmt.Println(fn(6))
}
