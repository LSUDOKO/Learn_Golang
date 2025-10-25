package main

import "fmt"

type vertex struct {
	lat, long float64
}

var a map[string]vertex

func main() {
	a = make(map[string]vertex)
	a["bells labs"] = vertex{
		40.14,
		45.25,
	}
	fmt.Println(a["bells labs"])
	m := make(map[string]int)
	m["name"] = 1
	m["roll_no"] = 22
	for key, val := range m {
		fmt.Println(key, val)
	}
}
