package main

import "fmt"
func pri[t any](items[] t ){// or [t interface{}]   if we want to add specifically int and string use [t int || string] 
	//[t comparable] particular type 
	for _,item := range items {
		fmt.Println(item)
	}
}


// // The function `pristring` takes a slice of strings as input and prints each string in the slice.
//func pristring(items[] string ){
// 	for _,item := range items {
// 		fmt.Println(item)
// 	}
// }
//LIFO
type stack[t any] struct{
	elements []t
}
func main(){
	nums:=[]int{1,2,3,4}
	numstr:=[]string {"arpit","soma","ram"}
	pri(nums)
	fmt.Println()
	pri(numstr)
	mystack:=stack[int]{
		elements: []int{1,2,3,4},

	}
	mystackstring:=stack[string]{
		elements: []string{"arpit","babu","john"},
	}
	fmt.Println(mystack)
	fmt.Println(mystackstring)
} 