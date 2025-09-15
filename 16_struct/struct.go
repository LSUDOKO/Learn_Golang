package main

import (
	"fmt"
)

// order
// type order struct {
// 	name      string
// 	roll      int
// 	marks     float32
// 	createdAt time.Time // nano second precesion
// }
// func newOrder(name string ,roll int, marks float32) *order{
// 	myorder:=order{
// 		name:name,
// 		roll:roll,
// 		marks:marks,
// 	}
// 	return &myorder
// }

// func (o *order) changeStaus(name string) {//pointer use to change staus
//
//		o.name =name
//	}
//
// func (o *order) getroll() int{// jab hame instances par change karna ho to pointer(*) ka istemal karna hota ha,agar hame sirf get karna ha to * ki jarurat nahi hoti
//
//		return o.roll
//	}
func main() {
	lang := struct {
		name  string
		truth bool
	}{"arpit", true}
	fmt.Println(lang.truth)

	// myorder:=newOrder("arpit",25,69.00)
	// fmt.Println(myorder)

	// if you are not set field then it set to be zero value of that datatype like 0,"",flase
	// myorder := order{
	// 	name: "arpit", roll: 25, marks: 50.25,
	// }
	// myorder.createdAt = time.Now()
	// myorder.changeStaus("ram")
	// fmt.Println("status change", myorder)
	// fmt.Println("get element",myorder.getroll())
	// fmt.Println("my order :", myorder)//not return nill in time part
	// fmt.Println(myorder.name)
	// 	myor2:=order{
	// 		name: "soma",
	// 		roll: 26,
	// 		marks: 65.25,
	// 		createdAt: time.Now(),
	// 	}
	// 	fmt.Println(myor2);
}
