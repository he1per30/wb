package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main18() {
	var result int64
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go increment(&result, &wg)
	}
	wg.Wait()

	fmt.Println(result)

}

// Инкрементим нашу переменную с помощью atomic
func increment(i *int64, wg *sync.WaitGroup) {
	atomic.AddInt64(i, 1)
	wg.Done()
}
