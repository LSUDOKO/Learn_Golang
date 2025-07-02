package main

import "fmt"

func add(nums ...int) int {
	count := 0
	for _, v := range nums {
		count = count + v
	}
	return count
}
func main() {
	// fmt.Println(1, 2, 3, 4, 5, "hello")
	nums := []int{1, 2, 3, 4, 5}
	result := add(nums...)
	fmt.Println(result)
	// result := add(3, 4, 5, 6, 7)
	// fmt.Println(result)
}
