package main

import (
	"fmt"
)

// func processnum(numchan chan int) {
// 	for nums:=range numchan{
// 		fmt.Println("processing number", nums)
// 		time.Sleep(time.Second)
// 	}

// }
// func sum(result chan int, num1 int, num2 int) {
// 	numresult := num1 + num2
// 	result <- numresult

// }

// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("Processing...")
// }

//	func email(emailchan chan string, done chan bool) {
//		defer func() { done <- true }()
//		for email := range emailchan {
//			fmt.Println("sending email to ", email)
//			time.Sleep(time.Second)
//		}
//	}
func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "arpit"
	}()
	for i := 0; i < 2; i++ {
		select {
		case chan1val := <-chan1:
			fmt.Println("recevied data from chan1", chan1val)
		case chan2val := <-chan2:
			fmt.Println("recevied data from chan2", chan2val)
		}
	} // emailchan := make(chan string, 100)
	// done := make(chan bool)
	// go email(emailchan,done)
	// for i := 0; i < 5; i++ {
	// 	emailchan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// fmt.Println("done sending")
	// //this is important
	// close(emailchan)
	// // emailchan<-"arpitsingh3202@gmail.com"
	// emailchan<- "adoranto737@gmail.com"

	// fmt.Println(<-emailchan)
	// fmt.Println(<-emailchan)
	//<-done
	// done := make(chan bool)
	// go task(done)
	// res := <-done // blocking both execute ho raha ha reciveing bhi and sending bhi
	// fmt.Println("recieved", res)
	// result := make(chan int)
	// go sum(result, 4, 5)
	// //time.Sleep(time.Second)
	// re := <-result//blocking
	// fmt.Println(re)

	// numchan := make(chan int)
	// go processnum(numchan)
	// for {
	// 	numchan <- rand.Intn(100)
	// }
	//numchan <- 12
	//time.Sleep(time.Second * 2)
	// message:=make(chan string)
	// message <- "arpit"//blocking
	// result:= <- message
	// fmt.Println(result)
}
