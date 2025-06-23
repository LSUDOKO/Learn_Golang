package main

import "fmt"

// import "time"
func main() {
	// simple switch
	// i:=22
	// switch i{
	// 	case 1:
	// 		fmt.Println("one")
	// 	case 2:
	// 		fmt.Println("two")
	// 	case 3:
	// 		fmt.Println("three")
	// 	default:
	// 		fmt.Println("default")
	// }

	// multiple condition switch
	// switch time.Now().Weekday(){
	// 	case time.Saturday,time.Sunday:
	// 		fmt.Println("weekend")
	// 	default:
	// 		fmt.Println("workday")
	// }
	// type switch
	whoami := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i am bool", t)
		case int:
			fmt.Println("i am int", t)
		default:
			fmt.Println("i am default", t)
		}
	}
	whoami(25)
}

// 1 hour 1 min
