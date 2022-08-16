package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main3() {
	var f int64
	vals := []int64{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	for i := range vals {
		wg.Add(1)
		go powerSum(wg, vals[i], &f)
	}
	wg.Wait()
	fmt.Println(f)
}

// Функция использует WaitGroup + atomic для конкурентной записи в переменную i
func powerSum(wg *sync.WaitGroup, val int64, i *int64) {
	atomic.AddInt64(i, val*val)
	wg.Done()
}
