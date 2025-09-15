package main
import "fmt"
type orderstatus int
const(
	Received orderstatus =iota
	Confirmed 
	Prepared
	Delivered
)
//enumerated types
func changestatus(status orderstatus){
	fmt.Println("changing order status to",status)
}
func main(){
	changestatus(Confirmed)
}
