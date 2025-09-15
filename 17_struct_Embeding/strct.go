package main

import "fmt"

type customer struct {
	name  string
	phone string
}
type order struct {
	name    string
	company string
	id      int
	customer
}

func main() {
	// cust := customer{
	// 	name:  "tpitt",
	// 	phone: "62358996675",
	// }
	myorder := order{
		name:    "arpit",
		company: "Microsoft",
		id:      12,
		// customer: cust,
		customer: customer{name: "john", phone: "6230555558"},
	}
	fmt.Println(myorder.customer)
}
