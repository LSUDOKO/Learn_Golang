package main

import "fmt"
// func yes(num int){
// 	fmt.Println(&num)//different memory address
// 	num =20
// 	fmt.Println(&num)//different memory address
// 	fmt.Println("the number is :",num)
// }
func yes(num *int){
	fmt.Println(&num)
	*num=5
	fmt.Println(num)
	fmt.Println(&num)

}
func main(){
	var num int =1
	// memory address
	fmt.Println(&num)
	yes(&num)
	//fmt.Println("the number is :",num)
}