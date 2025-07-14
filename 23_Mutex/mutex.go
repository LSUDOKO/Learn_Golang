package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu sync.Mutex
}

func (p *post) inc(wq *sync.WaitGroup) {
	defer wq.Done()
	p.mu.Lock()
	p.views++
	p.mu.Unlock()
}
func main() {
	var wq sync.WaitGroup
	mypost := post{views: 0}
	for i := 0; i < 100; i++ {
		wq.Add(1)
		go mypost.inc(&wq)

	}
	wq.Wait()
	fmt.Println(mypost.views)
}
