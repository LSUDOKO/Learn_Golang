package main
import "fmt"
import "maps"
func main(){
	// creating map
	m:=make(map[string]string)


	// setting an element in map
	m["name"]="Arpit"  
	fmt.Println(m)//print map[name:Arpit]

	// getting an element from map
	// fmt.Println(m["name"])
	fmt.Println(m["phone"])
	// imP:if key is not present in map then it will return zero value of that type
	// here it will return empty string

	// deleting an element from map
	// delete(m,"name")
	// clear(m)
	fmt.Println(m)
	fmt.Println(m["name"])
	// for checking element is present or not we can use comma ok syntax
	v,ok:=m["name"]//ok will return true or false and v value of map
	if ok{
		fmt.Println("present")
	}else{
		fmt.Println("not present")
	}
	fmt.Println(v,ok)
	// creating two equal map
	m1:=make(map[string]string)
	m1["name"]="Arpit"
	m1["phone"]="1234567890"
	m2:=make(map[string]string)
	m2["name"]="Arpit"
	m2["phone"]="1234567890"
	fmt.Println(maps.Equal(m1,m2))// for checking two maps are equal or not 
}