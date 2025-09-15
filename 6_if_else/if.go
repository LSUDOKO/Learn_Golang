package main
 import "fmt"
 func main(){
	// age :=12
	// we declare a variable in if condition
	if age:=12;age >=18{
		fmt.Println("you are eligible for voting")
	} else if age <18 && age >0 {
		fmt.Println("you are not eligible for voting")
	} else{
		fmt.Println("Please enter valid age")
	}
	// go does not have ternary operator
 }