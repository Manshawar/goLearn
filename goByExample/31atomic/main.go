package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64

	var wg sync.WaitGroup
	var num int64
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				num += 1
				atomic.AddUint64(&ops, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops)
	fmt.Println("num:", num)
}
