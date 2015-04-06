package main

import (
	"fmt"
	"sync"
	"time"
)

func f(n int) {
	fmt.Println("thread: ", n)
	time.Sleep(time.Second * 5)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 300; i++ {
		go f(i)
		wg.Add(1)
	}
	wg.Wait()
}
