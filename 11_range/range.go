package main

import "fmt"

// iterating over data structure using range
func main() {
	// num := []int{1, 2, 3, 4}
	// for i:=0;i<len(num);i++{
	// 	fmt.Println(num[i])
	// }
	// sum:=0
	// for _,v:=range(num){
	// 	sum=sum+v
	// }
	// fmt.Println(sum)
	// for i,v:=range(num){
	// 	fmt.Println(i,v)
	// }
	// iteration for map
	// m := map[string]string{"name": "Arpit", "phone": "6230603202"}
	// for i, v := range(m) {
	// 	fmt.Println(i, v)
	// }
	// itreation in string
	//300 --> 1byte, 2 byte
	num:="Arpit"
	for i,v:=range(num){
		fmt.Println(i,string(v))
	}

}
