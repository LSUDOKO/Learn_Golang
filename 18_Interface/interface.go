package main
import "fmt"

// interface method 
type paymenter interface{
	pay(amount float32)
}
type payment struct{
	gateway paymenter
}
func (p payment) makepayment(amount float32){
	// razorpaygw:=razorpay{}
	// razorpaygw.pay(amount)
	// stripepay:=stripe{}
	// stripepay.pay(amount)
	p.gateway.pay(amount);
}
type razorpay struct{}
func (r razorpay) pay(amount float32){
	//logic to make payment
	fmt.Println("making payment using razorpay",amount)
}

// type stripe struct{}
// func (s stripe) pay(amount float32){
// 	fmt.Println("make payment using stripe",amount)
// }

type fakepayment struct{}
func (f fakepayment) pay(amount float32){
fmt.Println("making payment using fake gateway for testing purpose ")	
}

type paypal struct{}
func (p paypal) pay(amount float32){
	fmt.Println("making payment using paypal is",amount)
}
func main(){
	// stipepay:=stripe{}
	// fakepa:=fakepayment{}
	// razorpayment:=razorpay{}
	paypalpa:=paypal{}
	newpayment:=payment{
		gateway: paypalpa,
	}
	newpayment.makepayment(100)
}